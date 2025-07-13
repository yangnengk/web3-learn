package data

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `uri:"username"`
	Password string `uri:"password"`
	Email    string `uri:"email"`
}

type Post struct {
	gorm.Model
	Author  User
	Title   string
	Content string
	UserID  uint
}

type Comment struct {
	gorm.Model
	Content string
	PostID  uint
	UserID  uint
}
