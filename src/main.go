package main

import (
	"os"
	"os/signal"
	"syscall"
	gw"gateway"
	"config"
	"path/filepath"
	"util"
    "log"
    "runtime"
    "web"
)

const DefaultConfigFile = "conf/default.json"

type EasyServer struct {
	proxyServer *gw.ProxyServer
	webServer   *web.WebServer
}

func CreateEasyServer() *EasyServer {
	return &EasyServer{proxyServer:new(gw.ProxyServer), webServer:new(web.WebServer)}
}

func (easyServer *EasyServer) Init(config *config.Config)  {
    easyServer.webServer.Init(config)
    easyServer.proxyServer.Init(config)
}

func (easyServer *EasyServer)Start()  {
    easyServer.webServer.Start()
    easyServer.proxyServer.Start()
}

func (easyServer *EasyServer)Stop()  {
    easyServer.proxyServer.Stop()
}

func (easyServer *EasyServer) CatchStopSignal() {
	sig := make(chan os.Signal)
	signal.Notify(sig, syscall.SIGKILL, syscall.SIGINT, syscall.SIGQUIT)
	go func() {
		log.Println("接受到退出信号：", <-sig)
		// 系统退出的资源保存和清理工作
		easyServer.Stop()
		log.Println("系统的清理工作已经完成")
	}()
}

func main() {
    homePath := util.HomePath()
    configValue, err := config.Load(filepath.Join(homePath, DefaultConfigFile))
    if err != nil {
        log.Println("加载配置错误")
    }
    runtime.GOMAXPROCS(runtime.NumCPU())
    easyServer := CreateEasyServer()
	easyServer.Init(configValue)
	easyServer.CatchStopSignal()
	easyServer.Start()
}
