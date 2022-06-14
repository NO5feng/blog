package views

import (
	"errors"
	"go_blog/common"
	"go_blog/service"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func (*HTMLApi) Category(w http.ResponseWriter, r *http.Request) {
	categoryTemplate := common.Template.Category //获取模板
	path := r.URL.Path                           // 获取地址 ——> /c/1
	cIdStr := strings.TrimPrefix(path, "/c/")    // strings.TrimPrefix函数的作用是：删除前缀
	cId, err := strconv.Atoi(cIdStr)             // 将字符串类型转换为int类型
	if err != nil {
		categoryTemplate.WriteError(w, errors.New("cIdStr转换int类型失败"))
		return
	}

	// r.ParseForm() 解析url传递的参数，对于post则解析相应包的主体（request body）
	if err := r.ParseForm(); err != nil {
		log.Println("表单获取失败：", err)
		categoryTemplate.WriteError(w, errors.New("系统错误，联系管理员"))
	}
	pageStr := r.Form.Get("page") // 取得表单中name为page的控件的值
	page := 1
	if pageStr == "" {
		pageStr = "1"
	}
	page, err = strconv.Atoi(pageStr) // 将字符串类型转换为int类型
	if err != nil {
		categoryTemplate.WriteError(w, errors.New("pageStr转换int类型失败"))
		return
	}
	// 每页显示的数量
	pageSize := 10
	categoryResponse, err := service.GetPostsByCategoryId(cId, page, pageSize)
	if err != nil {
		categoryTemplate.WriteError(w, err)
		return
	}
	categoryTemplate.WriteData(w, categoryResponse)
}
