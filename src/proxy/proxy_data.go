package proxy

import (
	"config"
	"structure"
	"sync"
)

type ProxyData struct {
	Service 		string
	Host 			string
	Port 			uint16
	Backends    	map[string]structure.Backend
	Deads 			map[string]structure.Backend
	ChannelManager 	*structure.ChannelManager
	mutex 			*sync.Mutex
}

func (data *ProxyData)Init(config *config.Config)  {
	data.Service = config.Service
	data.Host = config.Host
	data.Port = config.Port
	data.ChannelManager = new(structure.ChannelManager)
	data.ChannelManager.Init()
	data.setBackends(config.Backends)
	data.mutex = new(sync.Mutex)
}
func (data *ProxyData) setBackends(backends []structure.Backend) {
	data.Backends = make(map[string]structure.Backend)
	for _, backend := range backends {
		data.Backends[backend.Url()] = backend
	}
	data.Deads = make(map[string]structure.Backend)
}

