package main

import (
	"go_blog/common"
	"go_blog/router"
	"log"
	"net/http"
)

func init() {
	// 加载模板
	common.LoadTemplate()
}
func main() {
	// web程序，http协议 ip port
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	// 调用路由
	router.Router()
	if err := server.ListenAndServe(); err != nil {
		log.Println(err)
	}
}
