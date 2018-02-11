package main

import (
	"os"
	"os/signal"
	"syscall"
	gw"gateway"
	"config"
	"path/filepath"
	"log"
	"util"
)

const DefaultConfigFile = "conf/default.json"

type EasyServer struct {
	proxyServer *gw.ProxyServer
}

func CreateEasyServer() *EasyServer {
	return &EasyServer{proxyServer:new(gw.ProxyServer)}
}

func (easyServer *EasyServer) Init(config *config.Config)  {
	easyServer.proxyServer.Init(config)
}

func (easyServer *EasyServer)Start()  {
	easyServer.proxyServer.Start()
}

func (easyServer *EasyServer)Stop()  {
    easyServer.proxyServer.Stop()
}

func (easyServer *EasyServer) CatchStopSignal() {
	sig := make(chan os.Signal)
	signal.Notify(sig, syscall.SIGKILL, syscall.SIGINT, syscall.SIGQUIT)
	go func() {
		<-sig
		// 系统退出的资源保存和清理工作
		easyServer.Stop()
	}()
}
func main() {
	homePath := util.HomePath()
	configValue, err := config.Load(filepath.Join(homePath, DefaultConfigFile))
	if err != nil {
		log.Println()
	}
	easyServer := CreateEasyServer()
	easyServer.Init(configValue)
	easyServer.CatchStopSignal()
	easyServer.Start()
}
