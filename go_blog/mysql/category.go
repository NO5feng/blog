package mysql

import (
	"go_blog/models"
	"log"
)

func GetAllCategory() ([]models.Category, error) {
	rows, err := DB.Query("select * from blog_category")
	if err != nil {
		log.Println("GetAllCategory 查询错误：", err)
		return nil, err
	}
	var categors []models.Category
	for rows.Next() {
		var category models.Category
		err = rows.Scan(&category.Cid, &category.Name, &category.CreateAt, &category.UpdateAt)
		if err != nil {
			log.Println("GetAllCategory 取值错误：", err)
			return nil, err
		}
		categors = append(categors, category) // 向categors 数组中添加 category
	}
	return categors, nil
}
