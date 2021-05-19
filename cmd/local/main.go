package main

import (
	"fmt"
	"github.com/siriuswatch/myshawdowsockets"
	"github.com/siriuswatch/myshawdowsockets/cmd"
	"log"
	"net"
)

const (
	DefaultListenAddr = ":7448"
)

var version = "master"

func main() {
	log.SetFlags(log.Lshortfile)

	// 默认配置
	config := &cmd.Config{
		ListenAddr: DefaultListenAddr,
	}
	config.ReadConfig()
	config.SaveConfig()

	// 启动 local 端并监听
	lsLocal, err := myshawdowsockets.NewLsLocal(config.Password, config.ListenAddr, config.RemoteAddr)
	if err != nil {
		log.Fatalln(err)
	}
	log.Fatalln(lsLocal.Listen(func(listenAddr net.Addr) {
		log.Println(fmt.Sprintf(`
ss-local is working now: %s  
The config is:
	Local Listening Port:
	%s
	Remote Address:
	%s
	Password:
	%s`, version, listenAddr, config.RemoteAddr, config.Password))
	}))
}