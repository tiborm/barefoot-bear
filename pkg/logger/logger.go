package logger

import (
	"net/http"
	"runtime"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
)

var (
	once   sync.Once
	logger *Logger
)

type Logger struct {
	logger *logrus.Logger
}

func GetLogger() *Logger {
	once.Do(func() {
		logger = &Logger{}
		logger.logger = logrus.New()
		logger.logger.SetFormatter(&logrus.JSONFormatter{
			TimestampFormat: time.RFC3339,
		})
	})
	return logger
}

func (l Logger) Info(info string) {
	pc, _, line, ok := runtime.Caller(1)
	details := runtime.FuncForPC(pc)
	if ok && details != nil {
		l.logger.WithFields(logrus.Fields{
			"function": details.Name(),
			"line":     line,
		}).Info(info)
	} else {
		l.logger.Info(info)
	}
}

func (l Logger) Request(r *http.Request, info string) {
	pc, _, line, ok := runtime.Caller(1)
	details := runtime.FuncForPC(pc)
	if ok && details != nil {
		l.logger.WithFields(logrus.Fields{
			"requestId": r.Header.Get("X-Request-Id"),
			"method":   r.Method,
			"uri":      r.RequestURI,
			"function": details.Name(),
			"line":     line,
		}).Info(info)
	} else {
		l.logger.Info(info)
	}
}
