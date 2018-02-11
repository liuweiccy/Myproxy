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
	mutex 			*sync.RWMutex
}

func (data *ProxyData)Init(config *config.Config)  {
	data.Service = config.Service
	data.Host = config.Host
	data.Port = config.Port
	data.ChannelManager = new(structure.ChannelManager)
	data.ChannelManager.Init()
	data.setBackends(config.Backends)
	data.mutex = new(sync.RWMutex)
}

func (data *ProxyData) setBackends(backends []structure.Backend) {
	data.Backends = make(map[string]structure.Backend)
	for _, backend := range backends {
		data.Backends[backend.Url()] = backend
	}
	data.Deads = make(map[string]structure.Backend)
}

func (data *ProxyData) Start()  {

}

// 将清理存活的后台连接，存放到死亡的后台连接
func (data *ProxyData)cleanBackend(url string)  {
	data.mutex.Lock()
	defer data.mutex.Unlock()
	data.Deads[url] = data.Backends[url]
	delete(data.Backends, url)
}


func (data *ProxyData) cleanDeadend(url string)  {
	data.mutex.Lock()
	defer data.mutex.Unlock()
	data.Backends[url] = data.Deads[url]
	delete(data.Deads, url)
}

// 获取所有的后台连接地址
func (data *ProxyData) BackendUrls() []string {
	data.mutex.RLock()
	defer data.mutex.RUnlock()
	_map := data.Backends
	keys := make([]string, 0, len(_map))
	for k := range _map {
		keys = append(keys, k)
	}
	return keys
}

func (data *ProxyData) Clean() {
	cleanMap(data.Backends)
	cleanMap(data.Deads)
	data.ChannelManager.Clean()
}
func cleanMap(_map map[string]structure.Backend) {
    for k := range _map {
        delete(_map, k)
    }
}
