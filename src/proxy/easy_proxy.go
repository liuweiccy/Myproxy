package proxy

import (
	"config"
	"net"
)

type EasyProxy struct {
	data		*ProxyData
}

func (proxy *EasyProxy) Clean(url string) {
	panic("implement me")
}

func (proxy *EasyProxy) Recover(url string) {
	panic("implement me")
}

func (proxy *EasyProxy) Dispatch(con net.Conn) {
	panic("implement me")
}

func (proxy *EasyProxy) Close() {
	panic("implement me")
}

func (proxy *EasyProxy) Check()  {

}

func (proxy *EasyProxy)Init(config *config.Config)  {
	proxy.data = new(ProxyData)
	proxy.data.Init(config)
}

func (proxy *EasyProxy)Start()  {

}


