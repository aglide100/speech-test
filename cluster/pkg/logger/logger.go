package logger

import (
	"flag"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	debugMode = flag.Bool("debugMode", false, "using debug logger")
)

var log *zap.Logger
var atomicLevel = zap.NewAtomicLevel()

func init() {
	var err error

	flag.Parse()
	if (*debugMode) {
		log, err = zap.NewDevelopment()
		if err != nil {
			panic(err)
		}

		defer log.Sync()
		return
	}

	config := zap.NewProductionConfig()
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "timestamp"
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.StacktraceKey = ""            
	config.EncoderConfig = encoderConfig
 
	log, err = config.Build(zap.AddCallerSkip(1))
	if err != nil {
		panic(err)
	}
	
	defer log.Sync()
}

func SetLogLevel(level zapcore.Level) {
	atomicLevel.SetLevel(level)
}


func Info(message string, fields ...zap.Field) {
	// log.Sugar().Infow(message, fields)
	log.Info(message, fields...)
}

func Error(message string, fields ...zap.Field) {
	log.Error(message, fields...)
}

func Debug(message string, fields ...zap.Field) {
	log.Debug(message, fields...)
}