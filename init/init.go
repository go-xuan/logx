package init

import (
	"github.com/go-xuan/logx"

	"github.com/go-xuan/quanx/configx"
	"github.com/go-xuan/utilx/errorx"
	log "github.com/sirupsen/logrus"
)

func init() {
	// 设置默认日志输出
	log.SetOutput(logx.NewConsoleWriter())

	// 初始化日志配置
	errorx.Panic(Init())
}

func Init() error {
	if err := configx.LoadConfigurator(&logx.Config{}); err != nil {
		return errorx.Wrap(err, "init log failed")
	}
	return nil
}
