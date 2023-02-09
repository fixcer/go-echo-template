package logging

import (
	"fmt"
	configuration "go-backend-template/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"time"
)

var Log *zap.SugaredLogger

func Setup() {
	config := zap.NewProductionEncoderConfig()
	config.EncodeTime = zapcore.ISO8601TimeEncoder
	fileEncoder := zapcore.NewConsoleEncoder(config)
	writer := zapcore.AddSync(&lumberjack.Logger{
		Filename:   getLogFilePath() + getLogFileName(),
		MaxSize:    configuration.Cfg.App.Log.MaxSize, // megabytes
		MaxBackups: configuration.Cfg.App.Log.MaxBackup,
		MaxAge:     configuration.Cfg.App.Log.MaxAge, // days
	})
	defaultLogLevel := zapcore.DebugLevel
	core := zapcore.NewTee(
		zapcore.NewCore(fileEncoder, writer, defaultLogLevel),
		zapcore.NewCore(fileEncoder, zapcore.AddSync(os.Stdout), defaultLogLevel),
	)

	Log = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel)).Sugar()
}

func getLogFilePath() string {
	return fmt.Sprintf("%s%s", configuration.Cfg.App.RuntimeRootPath, configuration.Cfg.App.Log.SavePath)
}

func getLogFileName() string {
	return fmt.Sprintf("%s%s.%s",
		configuration.Cfg.App.Log.SaveName,
		time.Now().Format(configuration.Cfg.App.Log.TimeFormat),
		configuration.Cfg.App.Log.FileExt,
	)
}
