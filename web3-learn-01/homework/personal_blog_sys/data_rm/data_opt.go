package data_rm

import (
	"log"
	"time"

	"github.com/learn/personal_blog_sys/data_model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init() {
	// 初始化数据
	dsn := "root:123456@tcp(127.0.0.1:3306)/sqlx?charset=utf8mb4&parseTime=True&loc=Local"
	var error error
	db, error = gorm.Open(mysql.New(mysql.Config{
		DSN:                      dsn,
		DefaultStringSize:        256,
		DisableDatetimePrecision: true, // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
	}), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if error != nil {
		panic(error)
	}
	// 设置连接池
	if s, err := db.DB(); err != nil {
		panic(err)
	} else {
		s.SetMaxOpenConns(20)
		s.SetMaxIdleConns(10)
		s.SetConnMaxLifetime(time.Hour)
	}
	createTable()
}

func createTable() {
	err := db.AutoMigrate(&data_model.User{}, &data_model.Post{}, &data_model.Comment{})
	if err != nil {
		log.Println("创建表失败")
		panic(err)
	}
}

// CreateUser 创建用户
func CreateUser(user data_model.User) {
	db.Create(&user)
}

// GetUser 获取用户
func GetUser(username string) data_model.User {
	var user data_model.User
	db.Where("username = ?", username).First(&user)
	return user
}

// QueryUserInfo 查询用户信息,没有密码
func QueryUserInfo(username string) data_model.UserInfo {
	var user data_model.UserInfo
	db.Where("username = ?", username).First(&user)
	return user
}

// 创建文章
func CreatePost(post data_model.Post) {
	db.Create(&post)
}

func QueryPostList() []data_model.Post {
	var posts []data_model.Post
	db.Find(&posts)
	return posts
}

func UpdatePost(post data_model.Post) {
	db.Model(&post).Where("user_id = ?", post.UserID).Omit("CreatedAt", "UserID").Save(&post)
}

func DeletePost(post data_model.Post) {
	db.Debug().Where("user_id = ?", post.UserID).Delete(&data_model.Post{}, post.ID)
}

// 创建评论
func CreateComment(comment data_model.Comment) {
	db.Create(&comment)
}

// 查询评论
func QueryCommentList() []data_model.Comment {
	var comments []data_model.Comment
	db.Find(&comments)
	return comments
}
