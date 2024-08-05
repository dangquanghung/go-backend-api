package main

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	// 1.

	// 	sugar := zap.NewExample().Sugar()
	// 	sugar.Infof("Hello name:%s, age:%d", "dangquanghung", 40)

	// 	// logger
	// 	logger := zap.NewExample()
	// 	logger.Info("Hello", zap.String("Name", "dangquanghung"), zap.Int("age", 40))

	// 2.

	// logger := zap.NewExample()
	// logger.Info("Hello Example")

	// logger,_ = zap.NewDevelopment()
	// logger.Info("Hello NewDevelopment")

	// logger,_ = zap.NewProduction()
	// logger.Info("Hello NewProduction")

	// 3.

	encoder := getEncoderLog()
	sync := getWriterSync()
	core := zapcore.NewCore(encoder, sync, zapcore.InfoLevel)

	logger := zap.New(core, zap.AddCaller())

	logger.Info("Info Log", zap.Int("line", 1))
	logger.Error("Info Error", zap.Int("line", 2))

}

func getEncoderLog() zapcore.Encoder {

	encodeConfig := zap.NewProductionEncoderConfig()

	// 1722762361.3437698 -> 2024-8-4T16:39:01.877 + 0700
	encodeConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	// ts -> Time
	encodeConfig.TimeKey = "time"
	// from info INFO
	encodeConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	// "caller":"cli/main.log.go:23"

	encodeConfig.EncodeCaller = zapcore.ShortCallerEncoder //
	return zapcore.NewJSONEncoder(encodeConfig)
}

func getWriterSync() zapcore.WriteSyncer {

	if _, err := os.Stat("./log"); os.IsNotExist(err) {
		err := os.Mkdir("./log", 0755)
		if err != nil {
			zap.L().Fatal("Failed to create log directory", zap.Error(err))
		}
	}
	file, _ := os.OpenFile("./log/log.txt", os.O_CREATE|os.O_WRONLY, os.ModePerm)
	syncFile := zapcore.AddSync(file)
	syncConsole := zapcore.AddSync(os.Stderr)
	return zapcore.NewMultiWriteSyncer(syncConsole, syncFile)
}
