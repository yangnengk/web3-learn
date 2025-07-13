package data_model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
	Email    string `form:"email"`
}

type UserInfo struct {
	gorm.Model
	Username string `form:"username" binding:"required"`
	Email    string `form:"email"`
}

type Post struct {
	gorm.Model
	Title   string `form:"title" json:"title" binding:"required"`
	Content string `form:"content" json:"content" binding:"required"`
	UserID  uint
}

type Comment struct {
	gorm.Model
	Content string `form:"content" json:"content" binding:"required"`
	PostID  uint   `form:"postId" json:"postId" binding:"required"`
	UserID  uint
}

// TableName 设置表名
func (UserInfo) TableName() string {
	return "users"
}
