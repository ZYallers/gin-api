package logger

import (
	"code/app/cons"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"sync"
	"time"
)

var (
	loadedLoggers map[string]*zap.Logger
	loggersRWLock *sync.RWMutex
)

func init() {
	loadedLoggers = map[string]*zap.Logger{}
	loggersRWLock = new(sync.RWMutex)
}

func mkLogDir(dir string) {
	if _, err := os.Stat(dir); err != nil && os.IsNotExist(err) {
		if err = os.MkdirAll(dir, 0777); err != nil {
			panic(err)
		}
	}
}

func RouterLogger() *zap.Logger {
	mkLogDir(cons.LogDir)
	return getLogger(cons.LogDir+"/"+cons.Name, zap.DebugLevel)
}

func Debug(filename string, msg string, fields ...zap.Field) {
	var dir, fn string
	if filename == "" {
		dir = cons.LogDir
		fn = dir + "/" + cons.Name
	} else {
		dir = cons.LogDir + "/" + time.Now().Format("20060102")
		fn = dir + "/" + cons.Name + "_" + filename
	}
	mkLogDir(dir)
	getLogger(fn, zap.DebugLevel).Debug(msg, fields...)
}

func Info(filename string, msg string, fields ...zap.Field) {
	var dir, fn string
	if filename == "" {
		dir = cons.LogDir
		fn = dir + "/" + cons.Name
	} else {
		dir = cons.LogDir + "/" + time.Now().Format("20060102")
		fn = dir + "/" + cons.Name + "_" + filename
	}
	mkLogDir(dir)
	getLogger(fn, zap.InfoLevel).Info(msg, fields...)
}

func Error(filename string, msg string, fields ...zap.Field) {
	var dir, fn string
	if filename == "" {
		dir = cons.LogDir
		fn = dir + "/" + cons.Name
	} else {
		dir = cons.LogDir + "/" + time.Now().Format("20060102")
		fn = dir + "/" + cons.Name + "_" + filename
	}
	mkLogDir(dir)
	getLogger(fn, zap.ErrorLevel).Error(msg, fields...)
}

func getLogger(filename string, level zapcore.Level) *zap.Logger {
	loggersRWLock.RLock()
	if logger, ok := loadedLoggers[filename]; ok {
		loggersRWLock.RUnlock()
		return logger
	}
	loggersRWLock.RUnlock()

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString("[ZAP] " + t.Format("2006/01/02 15:04:05.000000"))
	}
	hook := lumberjack.Logger{Filename: filename + ".log", LocalTime: true}
	writeSyncer := zapcore.AddSync(&hook)
	core := zapcore.NewCore(zapcore.NewConsoleEncoder(encoderConfig), writeSyncer, level)
	logger := zap.New(core)

	loggersRWLock.Lock()
	loadedLoggers[filename] = logger
	loggersRWLock.Unlock()
	logger.Info("New ZapLogger Success", zap.String("filename", filename))

	return logger
}
