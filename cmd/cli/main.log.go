package main

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	encoder := getEncoderLog()
	sync := getLogWriter()
	core := zapcore.NewCore(encoder, sync, zapcore.InfoLevel)
	logger := zap.New(core, zap.AddCaller())
	defer logger.Sync()

	logger.Error("Failed to fetch URL.", zap.String("url", "http://example.com"), zap.Int("attempt", 3))
	logger.Info("Failed to fetch URL.", zap.String("url", "http://example.com"), zap.Int("attempt", 3))
}

func getLogWriter() zapcore.WriteSyncer {
	name := "./log/main.log"
	flag := os.O_WRONLY | os.O_CREATE | os.O_APPEND
	perm := os.ModePerm

	file, err := os.OpenFile(name, flag, perm)
	if err != nil {
		panic(err)
	}

	return zapcore.NewMultiWriteSyncer(zapcore.AddSync(file), zapcore.AddSync(os.Stderr))
}

func getEncoderLog() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.TimeKey = "timestamp"
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder

	return zapcore.NewJSONEncoder(encoderConfig)
}
