package api

import (
	"go_blog/common"
	"go_blog/service"
	"net/http"
)

func (*APIHandler) Login(w http.ResponseWriter, r *http.Request) {
	//接受用户名和密码 返回 对应的数据
	params := common.GetRequestJsonParam(r) //读取返回数据后，将数据赋值
	userName := params["username"].(string)
	passwd := params["passwd"].(string)
	loginRes, err := service.Login(userName, passwd)
	if err != nil {
		common.Error(w, err)
		return
	}
	common.Success(w, loginRes)
}
