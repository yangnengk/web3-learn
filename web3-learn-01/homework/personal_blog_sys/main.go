package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/learn/personal_blog_sys/log_opt"
	"github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
	"github.com/learn/personal_blog_sys/data_rm"
	"github.com/learn/personal_blog_sys/gin_opt"
	"github.com/learn/personal_blog_sys/jwt_util"
)

func main() {
	// gin.DisableConsoleColor()

	// 应用配置
	appConfig := &log_opt.AppConfig{
		Name:      "gin-app",
		Version:   "1.0.0",
		LogDir:    "./logs",
		MaxSizeMB: 100,
	}

	log_opt.InitLogger(appConfig)

	// 确保日志目录存在
	if err := os.MkdirAll(appConfig.LogDir, 0755); err != nil {
		log_opt.GetLogger().Fatal("Failed to create log directory:", err)
	}

	log_opt.GetLogger().Info("gin start")

	// gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	// 创建Gin引擎 为生产release模式
	gin.SetMode(gin.ReleaseMode)
	eg := gin.New()

	// 1. 重定向Gin默认日志到Logrus
	gin.DefaultWriter = &GinLogrusWriter{logger: log_opt.GetLogger()}
	gin.DefaultErrorWriter = &GinLogrusWriter{logger: log_opt.GetLogger()}

	// 2. 使用自定义中间件
	eg.Use(gin.Recovery())
	eg.Use(GinLoggerMiddleware(log_opt.GetLogger()))
	eg.Use(GinRecoveryMiddleware(log_opt.GetLogger()))

	eg.Use(func(ctx *gin.Context) {
		// 生产请求id
		requestID := ctx.GetHeader("X-Request-Id")
		if requestID == "" {
			requestID = generateRequestID()
		}

		ctx.Set("request_id", requestID)
		ctx.Writer.Header().Set("X-Request-Id", requestID)

		// 创建请求上下文日志记录器
		ctx.Set("logger", logrus.WithFields(logrus.Fields{
			"request_id": requestID,
			"client_ip":  ctx.ClientIP(),
			"method":     ctx.Request.Method,
			"path":       ctx.Request.URL.Path,
		}))
	})

	// 中间件
	eg.Use(HttpAop())
	eg.Use(ErrorHandler())

	data_rm.Init()

	eg.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello World",
		})
	})

	gin_opt.UserOperate(eg)
	gin_opt.PostOpt(eg)
	gin_opt.CommentOpt(eg)

	err := eg.Run(":8080")
	if err != nil {
		panic(err)
	}
}

func generateRequestID() string {
	return fmt.Sprintf("%d-%s", time.Now().UnixNano())
}

// 随机字符串生成
func randomString(n int) string {
	const letter = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, n)
	for i := range b {
		b[i] = letter[time.Now().UnixNano()%int64(len(letter))]
	}
	return string(b)
}

// GinLogrusWriter 实现 gin 的 io.Writer 接口
type GinLogrusWriter struct {
	logger *logrus.Logger
}

func (w *GinLogrusWriter) Write(p []byte) (n int, err error) {
	// 解析gin的日志格式
	msg := strings.TrimSpace(string(p))

	// 根据日志内容确定级别
	level := logrus.InfoLevel
	if strings.Contains(msg, "[ERROR]") {
		level = logrus.ErrorLevel
	} else if strings.Contains(msg, "[WARNING]") {
		level = logrus.WarnLevel
	}

	// 创建日志条目
	entry := w.logger.WithFields(logrus.Fields{
		"type": "gin_framework",
	})

	// 写入日志
	switch level {
	case logrus.InfoLevel:
		entry.Info(msg)
	case logrus.WarnLevel:
		entry.Warn(msg)
	case logrus.ErrorLevel:
		entry.Error(msg)
	default:
		entry.Info(msg)
	}
	return len(p), nil
}

// GinLoggerMiddleware ------------------gin hook---------------------
// GinLoggerMiddleware Gin 访问日志中间件
func GinLoggerMiddleware(logger *logrus.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 开始时间
		start := time.Now()

		// 请求路径
		path := ctx.Request.URL.Path
		query := ctx.Request.URL.RawQuery

		ctx.Next()

		// 结束时间
		end := time.Now()
		latency := end.Sub(start)

		// 客户端ip
		clientIp := ctx.ClientIP()

		method := ctx.Request.Method

		// 状态码
		statusCode := ctx.Writer.Status()

		// 错误信息
		errors := ctx.Errors.ByType(gin.ErrorTypePrivate).String()

		// 创建日志条目
		entry := logger.WithFields(logrus.Fields{
			"type":      "access",
			"status":    statusCode,
			"method":    method,
			"path":      path,
			"query":     query,
			"ip":        clientIp,
			"latency":   latency,
			"userAgent": ctx.Request.UserAgent(),
			"errors":    errors,
		})

		// 根据状态码设置日志级别
		if statusCode > 500 {
			entry.Error()
		} else if statusCode > 400 {
			entry.Warn()
		} else {
			entry.Info()
		}

	}
}

// GinRecoveryMiddleware Gin 恢复中间件
func GinRecoveryMiddleware(logger *logrus.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// 检查连接是否断开
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}
				// 记录错误日志
				logger.WithFields(logrus.Fields{
					"type":    "recovery",
					"error":   err,
					"request": ctx.Request.RequestURI,
					"stack":   getStack(),
				}).Error("Recovery from panic")

				if brokenPipe {
					ctx.Error(err.(error)) // 记录错误
					ctx.Abort()
				} else {
					ctx.AbortWithStatus(http.StatusInternalServerError)
				}
			}
		}()
		ctx.Next()
	}
}

// getStack 获取堆栈信息
func getStack() string {
	buf := make([]byte, 4096)
	n := runtime.Stack(buf, false)
	return string(buf[:n])
}

func HttpAop() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		t := time.Now()
		token := ctx.GetHeader("Authorization")
		url := ctx.FullPath()
		if url != "/ping" && url != "/user/login" && url != "/user/registry" {
			user, err := jwt_util.GetUserIDFromToken(token)
			if err != nil {
				ctx.JSON(http.StatusUnauthorized, gin.H{
					"error":   err.Error(),
					"status":  "false",
					"message": "token验证失败",
				})
				ctx.Abort()
				return
			}
			ctx.Set("user", user)
		}
		ctx.Next()
		cost := time.Since(t)
		ctx.Set("cost", cost)
		fmt.Printf("%s, 请求耗时: %v\n", url, cost)

		ctx.Writer.Status()
	}
}

func ErrorHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()
		if len(ctx.Errors) > 0 {
			err := ctx.Errors.Last().Err
			ctx.JSON(http.StatusBadRequest, gin.H{
				"status":  "false",
				"message": err.Error(),
			})
		}
	}
}
