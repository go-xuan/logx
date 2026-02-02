package logx

import (
	"github.com/go-xuan/configx"
	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetOutput(NewConsoleWriter()) // 设置默认日志输出
	Init()                            // 初始化日志配置
}

func Init() {
	logger := log.WithField("package", "logx")
	if err := configx.LoadConfigurator(&Config{}); err == nil {
		logger.Info("initialized success")
		return
	}
	logger.Warn("initialized failed")
}
