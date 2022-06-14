package mysql

import (
	"go_blog/models"
)

func GetPostPage(page, pagSize int) ([]models.Post, error) {
	page = (page - 1) * pagSize
	rows, err := DB.Query("select * from blog_post limit ?,?", page, pagSize)
	if err != nil {
		return nil, err
	}
	var posts []models.Post
	for rows.Next() {
		var post models.Post
		err := rows.Scan(
			&post.Pid,
			&post.Title,
			&post.Content,
			&post.Markdown,
			&post.CategoryId,
			&post.UserId,
			&post.ViewCount,
			&post.Type,
			&post.Slug,
			&post.CreateAt,
			&post.UpdateAt,
		)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func GetPostPageByCategoryId(cId, page, pagSize int) ([]models.Post, error) {
	page = (page - 1) * pagSize
	rows, err := DB.Query("select * from blog_post where category_id=? limit ?,?", cId, page, pagSize)
	if err != nil {
		return nil, err
	}
	var posts []models.Post
	for rows.Next() {
		var post models.Post
		err := rows.Scan(
			&post.Pid,
			&post.Title,
			&post.Content,
			&post.Markdown,
			&post.CategoryId,
			&post.UserId,
			&post.ViewCount,
			&post.Type,
			&post.Slug,
			&post.CreateAt,
			&post.UpdateAt,
		)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func CountGetAllPost() (count int) {
	rows := DB.QueryRow("select count(1) from blog_post") // select count(1)计算一个有多少符合条件的行
	_ = rows.Scan(&count)                                 // 从数据库驱动中获取一个值（该值会转换成src类型）,并将其存储到src,src满足driver.Value类型
	return
}

func CountGetAllPostByCategoryId(cId int) (count int) {
	rows := DB.QueryRow("select count(1) from blog_post where category_id=?", cId) // select count(1)计算一个有多少符合条件的行
	_ = rows.Scan(&count)                                                          // 从数据库驱动中获取一个值（该值会转换成src类型）,并将其存储到src,src满足driver.Value类型
	return
}

func GetPostById(pId int) (models.Post, error) {
	row := DB.QueryRow("select * from blog_post where pid=?", pId)
	var post models.Post
	if row.Err() != nil {
		return post, row.Err()
	}
	err := row.Scan(
		&post.Pid,
		&post.Title,
		&post.Content,
		&post.Markdown,
		&post.CategoryId,
		&post.UserId,
		&post.ViewCount,
		&post.Type,
		&post.Slug,
		&post.CreateAt,
		&post.UpdateAt,
	)
	if err != nil {
		return post, row.Err()
	}
	return post, nil
}
