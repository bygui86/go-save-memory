package logging

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/bygui86/go-save-memory/http-server/utils"
)

const (
	logEncodingEnvVar = "LOG_ENCODING"
	logLevelEnvVar    = "LOG_LEVEL"

	logEncodingDefault = "console"
	logLevelDefault    = "info"
)

var Log *zap.Logger
var SugaredLog *zap.SugaredLogger

func init() {
	encoding := utils.GetStringEnv(logEncodingEnvVar, logEncodingDefault)
	levelString := utils.GetStringEnv(logLevelEnvVar, logLevelDefault)
	level := zapcore.InfoLevel
	err := level.Set(levelString)
	if err != nil {
		panic(err)
	}
	buildLogger(encoding, level)
}

func buildLogger(encoding string, level zapcore.Level) {
	Log, _ = zap.Config{
		Encoding:         encoding,
		Level:            zap.NewAtomicLevelAt(level),
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
		EncoderConfig:    buildEncoderConfig(level),
	}.Build()
	SugaredLog = Log.Sugar()
}

func buildEncoderConfig(level zapcore.Level) zapcore.EncoderConfig {
	if level == zapcore.DebugLevel {
		return zapcore.EncoderConfig{
			MessageKey: "message",

			TimeKey:    "time",
			EncodeTime: zapcore.ISO8601TimeEncoder,

			LevelKey:    "level",
			EncodeLevel: zapcore.CapitalLevelEncoder,

			CallerKey:    "caller",
			EncodeCaller: zapcore.ShortCallerEncoder,
		}
	} else {
		return zapcore.EncoderConfig{
			MessageKey: "message",

			TimeKey:    "time",
			EncodeTime: zapcore.ISO8601TimeEncoder,

			LevelKey:    "level",
			EncodeLevel: zapcore.CapitalLevelEncoder,
		}
	}
}
