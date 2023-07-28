package main

import (
	"os"
	"time"

	"github.com/njfanxun/go-skeleton/cmd"
	"github.com/pterm/pterm"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func initZapLog() *zap.Logger {
	config := zap.NewProductionConfig()
	config.DisableStacktrace = true
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	config.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(time.DateTime + ".000")
	config.Encoding = "console"
	config.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	logger, _ := config.Build()
	return logger
}

func main() {
	logger := initZapLog()
	zap.ReplaceGlobals(logger)
	defer func(logger *zap.Logger) {
		_ = logger.Sync()
	}(logger)
	command := cmd.NewRootCommand()
	if err := command.Execute(); err != nil {
		pterm.Error.Printfln("程序执行错误:%s", err.Error())
		os.Exit(1)
	}
}
