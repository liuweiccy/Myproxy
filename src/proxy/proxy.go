package proxy

import "config"

type Proxy interface {
	Init(config *config.Config)
}

