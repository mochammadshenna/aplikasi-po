package logger

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/sirupsen/logrus"
)

func TestLevel(t *testing.T) {

	logger := logrus.New()
	logger.SetLevel(logrus.TraceLevel)

	fmt.Println(strings.ToUpper(logrus.DebugLevel.String()))

	logger.Trace("Tracing")
	logger.Debug("Debugging")
	logger.Info("Info")
	logger.Warn("Warning")
	logger.Warnf("Warning F")
	logger.Error("Error")
	// logger.Panic("Panic")
	// logger.Fatal("Fatal")
}

func TestLoggerOutput(t *testing.T) {
	logger := logrus.New()

	/**
	* Create output logger to destination file
	**/
	file, _ := os.OpenFile("logging.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	// if err != nil {
	// 	helper.PanicError(err)
	// }

	logger.SetOutput(file)

	logger.Info("This is a Info log message")
	logger.Warn("This is a Warning log message")
	logger.Error("This is a Error log message")
}

func TestLoggerFormatter(t *testing.T) {
	logger := logrus.New()

	/**
	* Create output logger to destination file by io.Writer method
	**/
	file, _ := os.OpenFile("logging.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	// if err != nil {
	// 	helper.PanicError(err)
	// }

	logger.SetOutput(file)
	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.Info("This is a Info log message")
	logger.Warn("This is a Warning log message")
	logger.Error("This is a Error log message")
}

func TestLoggerWithField(t *testing.T) {
	logger := logrus.New()

	/**
	* Create output logger to destination file using WithField by logrus
	**/
	file, _ := os.OpenFile("logging.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	// if err != nil {
	// 	helper.PanicError(err)
	// }

	logger.SetOutput(file)
	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.WithField("username", "Shenna").Info("This is a Info log message")
	logger.WithField("username", "Arza").Warn("This is a Warning log message")
	logger.WithField("username", "Sheira").Error("This is a Error log message")
}

func TestLoggerWithFields(t *testing.T) {
	logger := logrus.New()

	/**
	* Create output logger to destination file using WithFields []map[string]
	**/
	file, _ := os.OpenFile("logging.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	// if err != nil {
	// 	helper.PanicError(err)
	// }

	logger.SetOutput(file)
	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.WithFields(logrus.Fields{
		"name":     "Shenna",
		"password": "12h1uhjwdbmw",
	}).Infof("This is a With Fields")
}

func TestLoggerEntrySeeLog(t *testing.T) {
	logger := logrus.New()

	/**
	* Create output logger to destination file just to see log
	**/
	file, _ := os.OpenFile("logging.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	// if err != nil {
	// 	helper.PanicError(err)
	// }

	logger.SetOutput(file)
	logger.SetFormatter(&logrus.JSONFormatter{})

	entry := logrus.NewEntry(logger)

	entry.WithFields(logrus.Fields{
		"name":     "Arza",
		"password": "12h1uhjwdbmw",
	}).Infof("This is a With Fields")
}

type SampleHook struct {
}

func (s *SampleHook) Levels() []logrus.Level {
	return []logrus.Level{logrus.ErrorLevel, logrus.WarnLevel}
}

func (s *SampleHook) Fire(entry *logrus.Entry) error {
	fmt.Println("Sample Hook is called with level =>", entry.Level, entry.Message)
	return nil
}

func TestLogHook(t *testing.T) {
	logger := logrus.New()

	/**
	* Create output logger to destination file using Hook
	**/
	logger.AddHook(&SampleHook{})

	logger.Info("Testing hooks")
	logger.Warn("Testing hooks Warning")
}

func TestSingletonLog(t *testing.T) {
	logrus.Warn("TestSingletonLog Warning")
	logrus.Error("TestSingletonLog Error")

	/**
	* Create output logger to destination file using Singleton but becarefully this will be change globally
	**/
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.Warn("TestSingletonLog Warning")
	logrus.Error("TestSingletonLog Error")

}
