package logger

import (
    "fmt"
    "os"
    "path/filepath"
    "go.uber.org/zap"
    "go.uber.org/zap/zapcore"

    "FileServerWeb/config"
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

    // 创建目录，如果不存在的话
    err := os.MkdirAll(filepath.Dir(config.LOG_FILE_PATH), os.ModePerm)
    if err != nil {
        fmt.Println("Error creating directory:", err)
        return
    }

    file, err := os.Create(config.LOG_FILE_PATH)
    if err != nil {
        panic(err)
    }

    core := zapcore.NewCore(
        zapcore.NewJSONEncoder(encoderConfig),
        // zapcore.NewConsoleEncoder(encoderConfig),
        zapcore.AddSync(file),
        zap.NewAtomicLevelAt(zap.InfoLevel),
    )

    Logger = zap.New(core, zap.AddCaller())
}
