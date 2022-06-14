package service

import (
	"go_blog/config"
	"go_blog/models"
	"go_blog/mysql"
	"html/template"
)

func GetPostDetail(pId int) (*models.PostRes, error) {
	post, err := mysql.GetPostById(pId)
	if err != nil {
		return nil, err
	}
	categoryName := mysql.GetCategoryNameById(post.CategoryId)
	userName := mysql.GetUserNameById(post.UserId)
	postMore := models.PostMore{
		post.Pid,
		post.Title,
		post.Slug,
		template.HTML(post.Content), // 将post.Content的string类型转换为 HTML类型
		post.CategoryId,
		categoryName,
		post.UserId,
		userName,
		post.ViewCount,
		post.Type,
		models.DateDay(post.CreateAt),
		models.DateDay(post.UpdateAt),
	}
	var passRes = &models.PostRes{
		config.Cfg.Viewer,
		config.Cfg.System,
		postMore,
	}
	return passRes, nil
}
