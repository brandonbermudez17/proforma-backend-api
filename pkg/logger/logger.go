package logger

import (
	"os"

	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func getEncoderConfig() zapcore.EncoderConfig {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.MessageKey = "msg"
	encoderConfig.LevelKey = "level"
	encoderConfig.TimeKey = "time"
	encoderConfig.NameKey = "logger"
	encoderConfig.CallerKey = "file"
	encoderConfig.StacktraceKey = "stacktrace"
	encoderConfig.LineEnding = zapcore.DefaultLineEnding
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	encoderConfig.EncodeName = zapcore.FullNameEncoder

	return encoderConfig
}

func getJSONEncoder() zapcore.Encoder {
	config := getEncoderConfig()
	return zapcore.NewJSONEncoder(config)
}

func getConsoleEncoder() zapcore.Encoder {
	config := getEncoderConfig()
	return zapcore.NewConsoleEncoder(config)
}

func getEncoder(dev bool) zapcore.Encoder {
	if dev {
		return getConsoleEncoder()
	}
	return getJSONEncoder()
}

func getLogFile() *lumberjack.Logger {
	logpath := viper.GetString("log_path")
	file := new(lumberjack.Logger)
	file.Filename = logpath
	file.MaxSize = 128
	file.MaxAge = 7
	file.MaxBackups = 10
	file.Compress = false
	return file

}

func ProvideLogger() *zap.Logger {
	dev := viper.GetBool("development")
	hook := getLogFile()
	caller := zap.AddCaller()
	encoder := getEncoder(dev)

	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(zap.DebugLevel)

	var writes = []zapcore.WriteSyncer{zapcore.AddSync(hook)}

	// var mode zap.Option
	if dev {
		writes = append(writes, zapcore.AddSync(os.Stdout))
		// mode = zap.Development()
	}
	core := zapcore.NewCore(
		encoder,
		zapcore.NewMultiWriteSyncer(writes...),
		atomicLevel,
	)

	log := zap.New(core, caller)

	return log
}
