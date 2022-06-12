package models

import "go_blog/config"

type HomeResponse struct {
	config.Viewer // 配置查看器
	Categorys     []Category
	Posts         []PostMore
	Total         int // 总页数
	Page          int // 页面显示页数
	Pages         []int
	PageEnd       bool
}
