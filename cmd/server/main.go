package main

import (
	"fmt"
	"github.com/phayes/freeport"
	"github.com/siriuswatch/myshawdowsockets"
	"github.com/siriuswatch/myshawdowsockets/cmd"
	"log"
	"net"
	"os"
	"strconv"
)

var version = "master"

func main() {
	log.SetFlags(log.Lshortfile)

	// 优先从环境变量中获取监听端口
	port, err := strconv.Atoi(os.Getenv("LIGHTSOCKS_SERVER_PORT"))
	// 服务端监听端口随机生成
	if err != nil {
		port, err = freeport.GetFreePort()
	}
	if err != nil {
		// 随机端口失败就采用 7448
		port = 7448
	}
	// 默认配置
	config := &cmd.Config{
		ListenAddr: fmt.Sprintf(":%d", port),
		// 密码随机生成
		Password: myshawdowsockets.RandPassword(),
	}
	config.ReadConfig()
	config.SaveConfig()

	// 启动 server 端并监听
	lsServer, err := myshawdowsockets.NewLsServer(config.Password, config.ListenAddr)
	if err != nil {
		log.Fatalln(err)
	}
	log.Fatalln(lsServer.Listen(func(listenAddr net.Addr) {
		log.Println(fmt.Sprintf(`
ss-server is working now: %s 
The config is:
	Listening Port:
	%s
	Password:
	%s`, version, listenAddr, config.Password))
	}))
}