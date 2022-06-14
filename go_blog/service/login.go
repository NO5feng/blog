package service

import (
	"errors"
	"go_blog/models"
	"go_blog/mysql"
	"go_blog/utils"
)

func Login(userName, passwd string) (*models.LoginRes, error) {
	println(passwd)
	passwd = utils.Md5Crypt(passwd, "feng") //进行md5二次加密
	user := mysql.GetUser(userName, passwd)
	if user == nil {
		return nil, errors.New("账号密码错误")
	}
	uid := user.Uid
	//生成token jwt技术进行生成 令牌
	token, err := utils.Award(&uid)
	if err != nil {
		return nil, errors.New("token未能生成")
	}
	var userInfo models.UserInfo
	userInfo.Uid = user.Uid
	userInfo.UserName = user.UserName
	userInfo.Avatar = user.Avatar
	var lr = &models.LoginRes{
		token,
		userInfo,
	}
	return lr, nil
}
