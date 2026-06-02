package logx

import (
	"github.com/go-xuan/configx"
	"github.com/go-xuan/utilx/errorx"
	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetOutput(NewConsoleWriter()) // 设置默认日志输出
	_ = Initialize()                  // 初始化日志，先使用默认配置
}

func Initialize() error {
	logger := log.WithField("package", "ossx")
	if err := configx.LoadConfigurator(&Config{}); err == nil {
		logger.Info("initialize success")
		return nil
	}
	logger.Warn("initialize failed")
	return errorx.New("failed to initialize logx")
}
