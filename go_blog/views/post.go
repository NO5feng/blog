package views

import (
	"errors"
	"go_blog/common"
	"go_blog/service"
	"net/http"
	"strconv"
	"strings"
)

func (*HTMLApi) Detail(w http.ResponseWriter, r *http.Request) {
	detail := common.Template.Detail             // 获取模板
	path := r.URL.Path                           // 获取地址 ——> /p/1.html
	pIdStr := strings.TrimPrefix(path, "/p/")    // strings.TrimPrefix函数的作用是：删除前缀
	pIdStr = strings.TrimSuffix(pIdStr, ".html") // 删除后缀
	pId, err := strconv.Atoi(pIdStr)             // 将字符串类型转换为int类型
	if err != nil {
		detail.WriteError(w, errors.New("pIdStr转换int类型失败"))
		return
	}
	postRes, err := service.GetPostDetail(pId)
	if err != nil {
		detail.WriteError(w, errors.New("查询出错"))
		return
	}
	detail.WriteData(w, postRes)
}
