package logger

import (
	"application/app/constant"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"time"
)

var LoadedMap = map[string]*zap.Logger{}

func getNowTimeFormat() string {
	return time.Now().Format("20060102")
}

func mkdir() (string, error) {
	logDir := constant.LogDir + "/" + getNowTimeFormat()
	if _, err := os.Stat(logDir); err != nil && os.IsNotExist(err) {
		if err = os.MkdirAll(logDir, 0777); err != nil {
			return "", err
		}
	}
	return logDir, nil
}

func RouterLogger() *zap.Logger {
	if _, err := os.Stat(constant.LogDir); err != nil && os.IsNotExist(err) {
		if err = os.MkdirAll(constant.LogDir, 0777); err != nil {
			panic(err)
		}
	}
	return getLogger(constant.LogDir+"/"+constant.Name, zap.DebugLevel)
}

func Debug(filename string, msg string, fields ...zap.Field) {
	logDir, err := mkdir()
	if err != nil {
		panic(err)
	}
	getLogger(logDir+"/"+filename, zap.DebugLevel).Debug(msg, fields...)
}

func Info(filename string, msg string, fields ...zap.Field) {
	logDir, err := mkdir()
	if err != nil {
		panic(err)
	}
	getLogger(logDir+"/"+filename, zap.InfoLevel).Info(msg, fields...)
}

func Error(filename string, msg string, fields ...zap.Field) {
	logDir, err := mkdir()
	if err != nil {
		panic(err)
	}
	getLogger(logDir+"/"+filename, zap.ErrorLevel).Error(msg, fields...)
}

func getLogger(filename string, level zapcore.Level) *zap.Logger {
	if logger, ok := LoadedMap[filename]; ok {
		return logger
	}
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString("[zap-log] " + t.Format("2006/01/02 15:04:05.000000"))
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
	defer func() {
		logger.Info("New ZapLogger Success", zap.String("filename", filename))
		LoadedMap[filename] = logger
	}()
	return logger
}
