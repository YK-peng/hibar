package main

import (
	"flag"
	"hibar/cmd/comet/conf"
	"hibar/lib/logger"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func init() {
	// flag
	flag.StringVar(&conf.SocketPort, "socketport", ":7000", "comet server socket port")
	flag.StringVar(&conf.PProfPort, "pprofport", ":7001", "comet server PProf port")
	flag.StringVar(&conf.RpcPort, "rpcport", ":7005", "comet server rpc server ip")
	flag.StringVar(&conf.WebsocketPort, "websocketport", ":7006", "comet server websocket port")
	flag.Parse()
}

func main() {
	// log
	initLog()
	logger.LogAllFlags()

	// redis

	// rpc

	// socket

	// websocket

	// 注册服务

	// profile、metric

	// signal
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	signalSelect(c, onExit)
}

func initLog() {
	logger.SetLogger(logger.NewHiLogger())
}

func signalSelect(c chan os.Signal, onExit func()) {
	for {
		s := <-c
		//
		logger.Info("cometServer get signal", s)
		//
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			if nil != onExit {
				onExit()
			}
			time.Sleep(time.Second)
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}

}

func onExit() {
	// 退出
}
