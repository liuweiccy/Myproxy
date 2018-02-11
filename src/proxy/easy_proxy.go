package proxy

import (
	"config"
	"net"
	"log"
    "proxy/schedule"
    "time"
    "structure"
    "io"
)

const ChannelPairNum = 2

const DefaultTimeoutTime = 3

type EasyProxy struct {
	data		*ProxyData
	strategy    schedule.Strategy
}

func (proxy *EasyProxy) Clean(url string) {
	proxy.data.cleanBackend(url)
}

func (proxy *EasyProxy) Recover(url string) {
	proxy.data.cleanDeadend(url)
}

func (proxy *EasyProxy) Dispatch(con net.Conn) {
	if proxy.isBackendAvailable() {
		servers := proxy.data.BackendUrls()
	    url := proxy.strategy.Choose(con.RemoteAddr().String(), servers)
	    proxy.transfer(con, url)
	}
}

func (proxy *EasyProxy) Close() {
	proxy.data.Clean()
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

func (proxy *EasyProxy) transfer(conn net.Conn, remote string) {
   remoteConn, err := net.DialTimeout("tcp", remote, DefaultTimeoutTime * time.Second)
   if err != nil {
       conn.Close()
       // 链接失败，将放到死亡链接组，已被唤醒
       proxy.Clean(remote)
       log.Println("connect backend error:%s", err)
       return
   }
   sync := make(chan int, 1)
   channel := structure.Channel{SrcConn:conn, DstConn:remoteConn}
   // 存储通道
   go proxy.putChannel(&channel)
   // 交换数据
   go proxy.safeCopy(conn, remoteConn, sync)
   go proxy.safeCopy(remoteConn, conn, sync)
   // 删除已经交换数据的通道
   go proxy.closeChannel(&channel, sync)
}
func (proxy *EasyProxy) putChannel(channel *structure.Channel) {
    proxy.data.ChannelManager.PutChannel(channel)
}

func (proxy *EasyProxy) safeCopy(srcConn net.Conn, dstConn net.Conn, sync chan int) {
    io.Copy(srcConn, dstConn)
    defer srcConn.Close()
    sync <- 1
}

func (proxy *EasyProxy) closeChannel(channel *structure.Channel, sync chan int) {
    for i := 0; i < ChannelPairNum; i++ {
        <- sync
    }
    proxy.data.ChannelManager.DeleteChannel(channel)
}


