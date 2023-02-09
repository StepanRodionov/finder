package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

func New(level, name, version string) (*zap.Logger, error) {
	var logLevel zap.AtomicLevel
	if err := logLevel.UnmarshalText([]byte(level)); err != nil {
		return nil, err
	}

	encoder := zapcore.NewJSONEncoder(zapcore.EncoderConfig{
		TimeKey:        "timestamp",
		LevelKey:       "level",
		NameKey:        "facility",
		CallerKey:      "caller",
		FunctionKey:    "function",
		MessageKey:     "short_message",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.RFC3339TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	})

	logger := zap.New(
		zapcore.NewCore(encoder, os.Stdout, logLevel),
		zap.AddCaller(),
		zap.AddCallerSkip(1),
		zap.Fields(
			zap.String("release", version),
			zap.String("facility", name),
		),
	)

	return logger, nil
}
