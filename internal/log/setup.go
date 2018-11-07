package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.SugaredLogger

func init() {
	logger, _ := zap.Config{
		Level: zap.NewAtomicLevelAt(zap.DebugLevel),
		Development: false,
		Sampling: &zap.SamplingConfig{
			Initial: 100,
			Thereafter: 100,
		},
		Encoding: "json",
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey: "ts",
			LevelKey: "level",
			NameKey: "logger",
			CallerKey: "caller",
			MessageKey: "msg",
			StacktraceKey: "stacktrace",
			LineEnding: zapcore.DefaultLineEnding,
			EncodeLevel: zapcore.LowercaseLevelEncoder,
			EncodeTime: zapcore.ISO8601TimeEncoder,
			EncodeDuration: zapcore.SecondsDurationEncoder,
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
		OutputPaths: []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
	}.Build()

	Logger = logger.Sugar()
}
