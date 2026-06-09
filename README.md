# logx

日志系统，基于 logrus + lumberjack，支持日志级别、格式化输出与文件滚动。

## 安装

```bash
go get github.com/go-xuan/logx
```

## 快速开始

在 `conf/config.yaml` 中配置：

```yaml
log:
  level: "debug"
  format: "text"
  file:
    enable: true
    path: "./logs/app.log"
    maxSize: 100     # MB
    maxBackups: 10
    maxAge: 30       # 天
```

```go
import (
    "github.com/go-xuan/logx"
)

func main() {
    logx.Initialize()            // 从配置初始化
    logx.Info("server started")  // 直接使用全局 logger
}
```

## 主要功能

- **多级别日志** — Trace/Debug/Info/Warn/Error/Fatal/Panic
- **格式化输出** — text/json 格式
- **文件滚动** — 基于 lumberjack，支持按大小/数量/时间滚动
- **自动初始化** — 配合 configx 自动读取配置
