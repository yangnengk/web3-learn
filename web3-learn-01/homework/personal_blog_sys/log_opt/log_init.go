package log_opt

import (
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

var logger *logrus.Logger

func GetLogger() *logrus.Logger {
	return logger
}

func InitLogger(logPath string, maxSize int) *logrus.Logger {
	logger = logrus.New()
	// 设置日志为json格式
	logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})

	// 同时输出到文件和控制台
	logger.SetOutput(logrus.StandardLogger().Out)

	// 创建日期和大小轮转的hook
	logger.AddHook(newDateSizeRotateHook(logPath, maxSize))

	return logger
}

func newDateSizeRotateHook(logPath string, maxSize int) logrus.Hook {
	return &dateSizeRotateHook{
		logDir:    logPath,
		maxSize:   maxSize,
		writerMap: make(map[logrus.Level]*lumberjack.Logger),
	}
}

func (d *dateSizeRotateHook) Fire(entry *logrus.Entry) error {
	d.mu.Lock()
	defer d.mu.Unlock()

	currentDate := time.Now().Format("2006-01-02")

	if d.currentDay != currentDate {
		for _, w := range d.writerMap {
			w.Close()
		}
		d.writerMap = make(map[logrus.Level]*lumberjack.Logger)
		d.currentDay = currentDate
	}

	writer, exists := d.writerMap[entry.Level]
	if !exists {
		levelDir := filepath.Join(d.logDir, entry.Level.String())
		if err := os.MkdirAll(levelDir, 0755); err != nil {
			return err
		}

		logPath := filepath.Join(levelDir, currentDate+".log")
		writer = &lumberjack.Logger{
			Filename:   logPath,
			MaxSize:    d.maxSize,
			MaxBackups: 30,
			MaxAge:     30,
			Compress:   true,
		}
		d.writerMap[entry.Level] = writer
	}
	formatter := &logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	}
	line, err := formatter.Format(entry)
	if err != nil {
		return err
	}
	_, err = writer.Write(line)
	return err
}

func (d *dateSizeRotateHook) Levels() []logrus.Level {
	// 可以在这里控制哪些级别需要分离
	return []logrus.Level{
		logrus.PanicLevel,
		logrus.FatalLevel,
		logrus.ErrorLevel,
		logrus.WarnLevel,
		logrus.InfoLevel,
		// 可以注释掉不需要的级别
		// logrus.DebugLevel,
	}
}

type dateSizeRotateHook struct {
	logDir     string
	maxSize    int
	writerMap  map[logrus.Level]*lumberjack.Logger
	currentDay string
	mu         sync.Mutex
}
