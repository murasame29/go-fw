package main

import (
	"flag"
	"fmt"
	"murasame29/go-fw/src/frameworks/echo"
	"murasame29/go-fw/src/frameworks/gin"
	"murasame29/go-fw/src/frameworks/gorilla"
	"murasame29/go-fw/src/frameworks/net"
	"murasame29/go-fw/src/server"
)

var (
	FW string
)

func init() {
	flag.StringVar(&FW, "fw", "", "framework to use")
	flag.Parse()
}

func selectFramework(fw string) any {
	switch fw {
	case "gin":
		return gin.NewGinServer()
	case "echo":
		return echo.NewEchoServer()
	case "net":
		return net.NewNetHttpServer()
	case "gorilla":
		return gorilla.NewGorillaServer()
	default:
		panic(fmt.Sprintf("Unknown framework %s", fw))
	}
}

func main() {
	router := selectFramework(FW)
	server.NewServer().Run(router)
}
