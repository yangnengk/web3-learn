package main

import (
	"fmt"
	"github.com/learn/personal_blog_sys/log_opt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/learn/personal_blog_sys/data_rm"
	"github.com/learn/personal_blog_sys/gin_opt"
	"github.com/learn/personal_blog_sys/jwt_util"
)

func main() {
	// gin.DisableConsoleColor()

	log_opt.InitLogger("./log", 10000000)

	log_opt.GetLogger().Info("gin start")
	// f, _ := os.Create("gin.log")
	fileName := "blog-sys-gin-" + time.Now().Format("2006-01-02") + ".log"
	f, _ := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	eg := gin.New()

	eg.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		// 你的自定义格式
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))
	eg.Use(gin.Recovery())

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
