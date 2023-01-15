package log

import (
	"fmt"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Field = zap.Field

const (
	// prod uses info level logging
	Production = "prd"
	// staging uses info level logging
	Staging = "stg"
	// dev uses development logging
	Development = "dev"
)

var logger, _ = zap.NewProduction()
var sugar = logger.Sugar()

func Init(environment string) {
	switch environment {
	case Staging:
		fallthrough
	case Production:
		logger, _ = zap.NewProduction(zap.AddStacktrace(zapcore.ErrorLevel))
		Infof("log level set to info")
	case Development:
		fallthrough
	default:
		logger, _ = zap.NewDevelopment(zap.AddStacktrace(zapcore.WarnLevel))
		Infof("log level set to debug")
	}

	// we don't want to see that the top-level caller is log.* every time.
	//logger = logger.WithOptions(zap.AddCallerSkip(1))
	sugar = logger.Sugar()
}

// Debug only logs in Development mode
func Debug(msg string) {
	sugar.Debug(msg)
}

// Debugf is a helper function equivalent to Debug(fmt.Sprintf(msg, args))
func Debugf(msg string, args ...interface{}) {
	sugar.Debugf(msg, args...)
}

// Debugw is a helper function equivalent to Debug(fmt.Sprintf(msg, args))
// TODO this is less performant than zap Fields, but more readable,
func Debugw(msg string, fields ...interface{}) {
	sugar.Debugw(msg, fields...)
}

// Info logs a message with info log level.
func Info(msg string) {
	sugar.Info(msg)
}

// Infof logs a formatted message with info log level.
// Note that using this function is not recommended as it is less performant than using `Info` and passing a fieldList as structured context.
func Infof(msg string, args ...interface{}) {
	sugar.Infof(msg, args...)
}

// Infow logs a formatted message with info log level and a field list
func Infow(msg string, fields ...interface{}) {
	sugar.Infow(msg, fields...)
}

// Warn logs a message with warn log level.
func Warn(msg string) {
	sugar.Warn(msg)
}

// Warnf logs a formatted message with warn log level.
// Note that using this function is not recommended as it is less performant than using `Warn` and passing a fieldList as structured context.
func Warnf(msg string, args ...interface{}) {
	sugar.Warnf(msg, args...)
}

// Warnw logs a formatted message with Warn log level and a field list
func Warnw(msg string, err error, fields ...interface{}) {
	sugar.Warnw(fmt.Sprintf("%s : %v", msg, err), fields...)
}

// Error logs a simple error or string message with no structured context.
func Error(msg string) {
	sugar.Error(msg)
}

// Errorf logs a formatted message with error log level.
// Note that using this function is not recommended as it is less performant than using `Error` and passing a fieldList as structured context.
func Errorf(msg string, args ...interface{}) {
	sugar.Errorf(msg, args...)
}

// Errorw logs a formatted message with error log level.
// Note that using this function is not recommended as it is less performant than using `Error` and passing a fieldList as structured context.
func Errorw(msg string, err error, fields ...interface{}) {
	sugar.Errorw(fmt.Sprintf("%s : %v", msg, err), fields...)
}

// Fatal logs a message with fatal log level.
func Fatal(msg string) {
	sugar.Fatal(msg)
}

// Fatalf logs a formatted message with fatal log level.
// Note that using this function is not recommended as it is less performant than using `Fatal` and passing a fieldList as structured context.
func Fatalf(msg string, args ...interface{}) {
	sugar.Fatalf(msg, args...)
}
