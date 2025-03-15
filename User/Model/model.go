package models

import "gorm.io/gorm"

type UserBasic struct {
	gorm.Model
	UserName  string `json:"username"`
	PassWord  string `json:"passwordd"`
	Gender    int    `json:"gender"`
	Age       int    `json:"age"`
	Email     string `json:"email"`
	Telephone string `json:"telephone"`
}

func TableName(t *UserBasic) string {
	return "UserBasic"
}

// type ArticleBasic struct {
// 	gorm.Model
// 	Title    string `json:"title"`
// 	Subtitle string `json:"subtitle"`
// 	Labels   string `json:"labels"`
// 	Content  string `json:"content"`
// }

// func TableName(t *ArticleBasic) string {
// 	return "UserBasic"
// }
