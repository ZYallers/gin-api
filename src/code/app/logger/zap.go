package logger

import (
	"code/app/constant"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"time"
)

var LoadedMap = map[string]*zap.Logger{}

func mkLogDir(dir string) {
	if _, err := os.Stat(dir); err != nil && os.IsNotExist(err) {
		if err = os.MkdirAll(dir, 0777); err != nil {
			panic(err)
		}
	}
}

func RouterLogger() *zap.Logger {
	mkLogDir(constant.LogDir)
	return getLogger(constant.LogDir+"/"+constant.Name, zap.DebugLevel)
}

func Debug(filename string, msg string, fields ...zap.Field) {
	var dir, fn string
	if filename == "" {
		dir = constant.LogDir
		fn = dir + "/" + constant.Name
	} else {
		dir = constant.LogDir + "/" + time.Now().Format("20060102")
		fn = dir + "/" + filename
	}
	mkLogDir(dir)
	getLogger(fn, zap.DebugLevel).Debug(msg, fields...)
}

func Info(filename string, msg string, fields ...zap.Field) {
	var dir, fn string
	if filename == "" {
		dir = constant.LogDir
		fn = dir + "/" + constant.Name
	} else {
		dir = constant.LogDir + "/" + time.Now().Format("20060102")
		fn = dir + "/" + filename
	}
	mkLogDir(dir)
	getLogger(fn, zap.InfoLevel).Info(msg, fields...)
}

func Error(filename string, msg string, fields ...zap.Field) {
	var dir, fn string
	if filename == "" {
		dir = constant.LogDir
		fn = dir + "/" + constant.Name
	} else {
		dir = constant.LogDir + "/" + time.Now().Format("20060102")
		fn = dir + "/" + filename
	}
	mkLogDir(dir)
	getLogger(fn, zap.ErrorLevel).Error(msg, fields...)
}

func getLogger(filename string, level zapcore.Level) *zap.Logger {
	if logger, ok := LoadedMap[filename]; ok {
		return logger
	}
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString("[ZAP] " + t.Format("2006/01/02 15:04:05.000000"))
	}
	hook := lumberjack.Logger{
		Filename:   filename + ".log",
		MaxSize:    100,
		MaxBackups: 30,
		MaxAge:     30,
		LocalTime:  true,
		Compress:   false,
	}
	writeSyncer := zapcore.AddSync(&hook)
	core := zapcore.NewCore(zapcore.NewConsoleEncoder(encoderConfig), writeSyncer, level)
	logger := zap.New(core)
	LoadedMap[filename] = logger
	logger.Info("New ZapLogger Success", zap.String("filename", filename))
	return logger
}
