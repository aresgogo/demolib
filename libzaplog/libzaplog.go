package libzaplog

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func ZapInitial() (*zap.Logger, error) {
	var conf = zap.NewProductionConfig()
	var logger *zap.Logger
	var err error
	conf = zap.NewProductionConfig()
	conf.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-12 15:04:05")
	logger, err = conf.Build()
	if nil != err {
		return nil, err
	}
	return logger, nil
}
