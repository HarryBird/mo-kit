package zap

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func DefaultEncoder() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		TimeKey:        "ts",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stack",
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
	}
}

func DefaultLevel() zap.AtomicLevel {
	return zap.NewAtomicLevelAt(zapcore.InfoLevel)
}

func DefaultSyncer() []zapcore.WriteSyncer {
	return []zapcore.WriteSyncer{zapcore.AddSync(os.Stdout)}
}

func DefaultCore() zapcore.Core {
	return zapcore.NewCore(
		zapcore.NewJSONEncoder(DefaultEncoder()),
		zapcore.NewMultiWriteSyncer(DefaultSyncer()...),
		DefaultLevel(),
	)
}

func DevelopmentLogger() *zap.Logger {
	return zap.New(
		DefaultCore(),
		zap.AddStacktrace(
			zap.NewAtomicLevelAt(zapcore.ErrorLevel)),
		zap.AddCaller(),
		zap.AddCallerSkip(3),
		zap.Development(),
	)
}
