package web

import (
    "config"
    "util"
    "net/http"
    _ "net/http/pprof"
    "log"
)

type WebServer struct {
    host string
    port uint16
}

func (webServer *WebServer)Init(config *config.Config)  {
    webServer.host = config.Host
    webServer.port = config.WebPort
}

func (webServer *WebServer)Start()  {
    go func() {
        webServer.AddHandler()
        url := util.HostAndPortToAddress(webServer.host, webServer.port)
        err := http.ListenAndServe(url, nil)
        if err != nil {
            log.Println("Create web server error: ", err)
        }
    }()
}

func (webServer *WebServer)AddHandler()  {
    http.HandleFunc(StatisticsUrl, Statistic)
}



