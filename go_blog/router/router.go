package router

import (
	"go_blog/views"
	"net/http"
)

func Router() {
	//1.页面 views 2. api 数据（json) 3.静态资源
	http.HandleFunc("/", views.HTML.Index)
	//http.HandleFunc("/", api.API.SaveAndUpdatePost)
	// 将resource 映射到 public/resource 否则将出现错乱
	// http.Handle 就是一个普通函数func(http.ResponseWriter,*http.Request)满足Handle接口的适配器
	http.Handle("/resource/", http.StripPrefix("/resource", http.FileServer(http.Dir("public/resource"))))
}
