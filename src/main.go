package main

import (
	"os"
	"os/signal"
	"syscall"
)

type EasyServer struct {

}

func (easyServer *EasyServer) CatchStopSignal() {
	sig := make(chan os.Signal)
	signal.Notify(sig, syscall.SIGKILL, syscall.SIGINT, syscall.SIGQUIT)
	go func() {
		<-sig
		// TODO 系统退出的资源保存和清理工作

	}()
}

func main()  {

}
