package data_rm

import (
	"github.com/learn/personal_blog_sys/data_model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
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
