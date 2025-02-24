package main

import (
	"gin-chat-demo/conf"
	"gin-chat-demo/router"
)

func main() {
	conf.Init()
	r := router.NewRouter()
	_ = r.Run(conf.HttpPort)
}
