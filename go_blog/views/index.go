package views

import (
	"errors"
	"go_blog/common"
	"go_blog/service"
	"log"
	"net/http"
	"strconv"
)

func (*HTMLApi) Index(w http.ResponseWriter, r *http.Request) {
	// Request 请求是为获取此响应而发送的请求。请求的主体为零（已被消耗）这仅为客户端请求填充
	index := common.Template.Index // 拿到模板
	// 数据库查询
	if err := r.ParseForm(); err != nil {
		log.Println("表单获取失败：", err)
		index.WriteError(w, errors.New("系统错误，联系管理员"))
	}
	pageStr := r.Form.Get("page")
	page := 1
	if pageStr != "" {
		page, _ = strconv.Atoi(pageStr)
	}

	// 每页显示的数量
	pageSize := 10
	hr, err := service.GetAllIndexInfo(page, pageSize)
	if err != nil {
		log.Println("views/Index获取失败：", err)
		index.WriteError(w, errors.New("系统错误，联系管理员"))
	}
	index.WriteData(w, hr)

}
