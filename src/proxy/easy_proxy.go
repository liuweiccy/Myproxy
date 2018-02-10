package proxy

import "config"

type EasyProxy struct {
	data		*ProxyData
}

func (proxy EasyProxy)Init(config *config.Config)  {
	proxy.data = new(ProxyData)
	proxy.data.Init(config)
}

