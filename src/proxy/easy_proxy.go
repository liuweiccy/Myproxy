package proxy

import (
	"config"
	"net"
	"log"
)

type EasyProxy struct {
	data		*ProxyData
}

func (proxy *EasyProxy) Clean(url string) {
	proxy.data.cleanBackend(url)
}

func (proxy *EasyProxy) Recover(url string) {
	proxy.data.cleanDeadend(url)
}

func (proxy *EasyProxy) Dispatch(con net.Conn) {
	if proxy.isBackendAvailable() {
		//servers := proxy.data.
	}
}

func (proxy *EasyProxy) Close() {
	panic("implement me")
}

func (proxy *EasyProxy) Check()  {
	for _, backend := range proxy.data.Backends {
		_, err := net.Dial("tcp", backend.Url())
		if err != nil {
			proxy.Clean(backend.Url())
		} else {
			log.Println(backend.Url() + " is keeplive")
		}
	}

	for _, deadend := range proxy.data.Deads {
		_, err := net.Dial("tcp", deadend.Url())
		if err == nil {
			proxy.Recover(deadend.Url())
			log.Println(deadend.Url(), "is recover connect")
		}
	}
}

func (proxy *EasyProxy)Init(config *config.Config)  {
	proxy.data = new(ProxyData)
	proxy.data.Init(config)
}

func (proxy *EasyProxy)Start()  {

}
func (proxy *EasyProxy) isBackendAvailable() bool {
	return len(proxy.data.Backends) > 0
}


