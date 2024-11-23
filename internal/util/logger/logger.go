package logger

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"

	stackdriver "github.com/TV4/logrus-stackdriver-formatter"
	"github.com/mochammadshenna/aplikasi-po/internal/state"

	"github.com/natefinch/lumberjack"
	"github.com/sirupsen/logrus"
)

var Logger = logrus.New()

var lf loggerField

func Init() {
	lf = newLoggerField()

	Logger.SetFormatter(stackdriver.NewFormatter(
		stackdriver.WithStackSkip("github.com/mochammadshenna/aplikasi-po/util/logger")),
	)

	logPath := "util/logger/logging.log"

	// output to file
	lumberjackLogger := &lumberjack.Logger{
		Filename:   filepath.ToSlash(logPath),
		MaxSize:    500,   // MB
		MaxBackups: 100,   // MB
		MaxAge:     14,    // days
		Compress:   false, // disabled by default
	}

	lumberjackLogger.Write([]byte(fmt.Sprintf("current time:%v\n", time.Now())))

	mw := io.MultiWriter(os.Stdout, lumberjackLogger)
	Logger.SetOutput(mw)

	/**
	* If the logger save on file and watcher it
	**/
	// logName := filepath.Base(logPath)
	// watcher, err := fsnotify.NewWatcher()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer watcher.Close()

	// go func() {
	// 	for event := range watcher.Events {
	// 		if event.Op&fsnotify.Remove == fsnotify.Remove &&
	// 			event.Name == logName {
	// 			log.Println("rotate log", event.Name)
	// 			lumberjackLogger.Rotate()
	// 		}
	// 	}
	// }()

	// err = watcher.Add(filepath.Dir(logPath))
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// file, err := os.OpenFile("util/logger/logging.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	// if err != nil {
	// 	panic(err)
	// }

	// Logger.SetOutput(file)
}

type loggerField struct {
	// Custom field
	RequestId     string `json:"requestId"`
	Latency       string `json:"latency"`
	RequestMethod string `json:"requestMethod"`
	Resource      string `json:"resource"`
	UserAgent     string `json:"userAgent"`
	PlatformType  string `json:"platformType"`
	Platform      string `json:"platform"`
	Version       string `json:"version"`
	Status        string `json:"status"`
	XForwardedFor string `json:"xForwardedFor"`

	// Field handle by logger
	Message        string `json:"message"`
	Severity       string `json:"severity"`
	Timestamp      string `json:"timestamp"`
	SourceLocation string `json:"sourceLocation"`
}

func newLoggerField() loggerField {
	return loggerField{
		RequestId:      "requestId",
		Latency:        "latency",
		RequestMethod:  "requestMethod",
		Resource:       "resource",
		UserAgent:      "userAgent",
		PlatformType:   "platformType",
		Platform:       "platform",
		Version:        "version",
		Status:         "status",
		XForwardedFor:  "xForwardedFor",
		Message:        "message",
		Severity:       "severity",
		Timestamp:      "timestamp",
		SourceLocation: "sourceLocation",
	}
}

func LoggerField() loggerField {
	return lf
}

func Trace(ctx context.Context, args ...interface{}) {
	fields := withFields(ctx)
	if len(fields) > 0 {
		Logger.WithFields(fields).Trace(args...)
		return
	}
	Logger.Trace(args...)
}

func Tracef(ctx context.Context, format string, args ...interface{}) {
	fields := withFields(ctx)
	if len(fields) > 0 {
		Logger.WithFields(fields).Tracef(format, args...)
		return
	}
	Logger.Tracef(format, args...)
}

func Debug(ctx context.Context, args ...interface{}) {
	fields := withFields(ctx)
	if len(fields) > 0 {
		Logger.WithFields(fields).Debug(args...)
		return
	}
	Logger.Debug(args...)
}

func Debugf(ctx context.Context, format string, args ...interface{}) {
	fields := withFields(ctx)
	if len(fields) > 0 {
		Logger.WithFields(fields).Debugf(format, args...)
		return
	}
	Logger.Debugf(format, args...)
}

func Info(ctx context.Context, args ...interface{}) {
	fields := withFields(ctx)
	if len(fields) > 0 {
		Logger.WithFields(fields).Info(args...)
		return
	}
	Logger.Info(args...)
}

func Infof(ctx context.Context, format string, args ...interface{}) {
	fields := withFields(ctx)
	if len(fields) > 0 {
		Logger.WithFields(fields).Infof(format, args...)
		return
	}
	Logger.Infof(format, args...)
}

func Warn(ctx context.Context, args ...interface{}) {
	fields := withFields(ctx)
	if len(fields) > 0 {
		Logger.WithFields(fields).Warn(args...)
		return
	}
	Logger.Warn(args...)
}

func Warnf(ctx context.Context, format string, args ...interface{}) {
	fields := withFields(ctx)
	if len(fields) > 0 {
		Logger.WithFields(fields).Warnf(format, args...)
		return
	}
	Logger.Warnf(format, args...)
}

func Error(ctx context.Context, args ...interface{}) {
	fields := withFields(ctx)
	if len(fields) > 0 {
		Logger.WithFields(fields).Error(args...)
		return
	}
	Logger.Error(args...)
}

func Errorf(ctx context.Context, format string, args ...interface{}) {
	fields := withFields(ctx)
	if len(fields) > 0 {
		Logger.WithFields(fields).Errorf(format, args...)
		return
	}
	Logger.Errorf(format, args...)
}

func Fatal(ctx context.Context, args ...interface{}) {
	fields := withFields(ctx)
	if len(fields) > 0 {
		Logger.WithFields(fields).Fatal(args...)
		return
	}
	Logger.Fatal(args...)
}

func Fatalf(ctx context.Context, format string, args ...interface{}) {
	fields := withFields(ctx)
	if len(fields) > 0 {
		Logger.WithFields(fields).Fatalf(format, args...)
		return
	}
	Logger.Fatalf(format, args...)
}

func Panic(ctx context.Context, args ...interface{}) {
	fields := withFields(ctx)
	if len(fields) > 0 {
		Logger.WithFields(fields).Panic(args...)
		return
	}
	Logger.Panic(args...)
}

func Panicf(ctx context.Context, format string, args ...interface{}) {
	fields := withFields(ctx)
	if len(fields) > 0 {
		Logger.WithFields(fields).Panicf(format, args...)
		return
	}
	Logger.Panicf(format, args...)
}

func withFields(ctx context.Context) logrus.Fields {
	fields := logrus.Fields{}

	requestId := ctx.Value(state.HttpHeaders().RequestId)
	if requestId != "" {
		fields[LoggerField().RequestId] = requestId
	}

	platformType := ctx.Value(state.HttpHeaders().PlatformType)
	if platformType != "" {
		fields[LoggerField().PlatformType] = platformType
	}

	platform := ctx.Value(state.HttpHeaders().Platform)
	if platform != "" {
		fields[LoggerField().Platform] = platform
	}

	version := ctx.Value(state.HttpHeaders().Version)
	if version != "" {
		fields[LoggerField().Version] = version
	}
	return fields
}
