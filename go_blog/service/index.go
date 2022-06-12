package service

import (
	"go_blog/config"
	"go_blog/models"
	"go_blog/mysql"
)

func GetAllIndexInfo() (*models.HomeResponse, error) {
	categorys, err := mysql.GetAllCategory()
	if err != nil {
		return nil, err
	}
	var posts = []models.PostMore{
		{
			Pid:          1,
			Title:        "go博客",
			Content:      "内容",
			UserName:     "海峰",
			ViewCount:    123,
			CreateAt:     "2022-06-06",
			CategoryId:   1,
			CategoryName: "go",
			Type:         0,
		},
	}
	var hr = &models.HomeResponse{
		config.Cfg.Viewer,
		categorys,
		posts,
		1,
		1,
		[]int{1},
		true,
	}
	return hr, err
}
