package logger

import (
	"os"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.Logger

func init() {
    encoderConfig := zapcore.EncoderConfig{
        MessageKey:   "msg",
        LevelKey:     "lv",
        TimeKey:      "time",
        CallerKey:    "caller",
        EncodeLevel:  zapcore.CapitalLevelEncoder,
        EncodeTime:   zapcore.ISO8601TimeEncoder,
        EncodeCaller: zapcore.ShortCallerEncoder,
    }

    file, err := os.Create("myapp.log")
    if err != nil {
        panic(err)
    }

    core := zapcore.NewCore(
        zapcore.NewJSONEncoder(encoderConfig),
        zapcore.AddSync(file),
        zap.NewAtomicLevelAt(zap.InfoLevel),
    )

    Logger = zap.New(core, zap.AddCaller())
}
