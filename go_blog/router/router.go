package router

import (
	"go_blog/api"
	"go_blog/views"
	"net/http"
)

func Router() {
	//1.页面 views 2. api 数据（json) 3.静态资源
	//Go中HTTP通信时，客户端请求信息封装在http.Request对象中，服务端返回的响应报文会被保存在http.Response结构体中。
	//需要注意的是，发送给客户端响应的并不是http.Response，而是通过http.ResponseWriter接口来实现的。
	http.HandleFunc("/", views.HTML.Index)
	//http://localhost:8080/c/1 1参数 分类的id
	http.HandleFunc("/c/", views.HTML.Category)
	http.HandleFunc("/login", views.HTML.Login) // 登入接口
	http.HandleFunc("/api/v1/login", api.API.Login)
	http.HandleFunc("/p/", views.HTML.Detail)
	//http.HandleFunc("/api/v1/post", api.API.SaveAndUpdatePost)
	// 将resource 映射到 public/resource 否则将出现错乱
	// http.Handle 就是一个普通函数func(http.ResponseWriter,*http.Request)满足Handle接口的适配器
	http.Handle("/resource/", http.StripPrefix("/resource", http.FileServer(http.Dir("public/resource"))))
}
