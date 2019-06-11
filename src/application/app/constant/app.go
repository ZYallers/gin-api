package constant

import (
	"time"
)

const (
	Name                  = "ginapi.zyall.com"
	Version               = "1.1.0"
	HttpServerAddr        = "127.0.0.1:8087"
	HttpServerIdleTimeout = 30 * time.Second
	LogDir                = "./log"
	DbDebug               = false
)
