package service

import (
	"go_blog/config"
	"go_blog/models"
	"go_blog/mysql"
	"html/template"
)

func GetPostsByCategoryId(cId, page, pageSize int) (*models.CategoryResponse, error) {
	// 获取 个人信息页面标签内容
	categorys, err := mysql.GetAllCategory()
	if err != nil {
		return nil, err
	}

	// 获取post页面内容（文章内容）
	posts, err := mysql.GetPostPageByCategoryId(cId, page, pageSize)
	var postMores []models.PostMore
	for _, post := range posts {
		categoryName := mysql.GetCategoryNameById(post.CategoryId)
		userName := mysql.GetUserNameById(post.UserId)
		// 将 post.Content 进行切割分页
		content := []rune(post.Content)
		if len(content) > 100 {
			content = content[0:100]
		}
		postMore := models.PostMore{
			post.Pid,
			post.Title,
			post.Slug,
			template.HTML(content), // 将post.Content的string类型转换为 HTML类型
			post.CategoryId,
			categoryName,
			post.UserId,
			userName,
			post.ViewCount,
			post.Type,
			models.DateDay(post.CreateAt),
			models.DateDay(post.UpdateAt),
		}
		postMores = append(postMores, postMore)
	}

	total := mysql.CountGetAllPostByCategoryId(cId)
	pagesCount := (total-1)/10 + 1
	var pages []int
	for i := 0; i < pagesCount; i++ {
		pages = append(pages, i+1)
	}
	var hr = &models.HomeResponse{
		config.Cfg.Viewer,
		categorys,
		postMores,
		total,
		page,
		pages,
		page != pagesCount,
	}

	categoryName := mysql.GetCategoryNameById(cId)
	categoryResponse := &models.CategoryResponse{
		hr,
		categoryName,
	}
	return categoryResponse, nil
}
