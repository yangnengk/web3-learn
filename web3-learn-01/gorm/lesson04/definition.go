package lesson04

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"strings"
)

// 自定义类型
type strArr []string
type UserD struct {
	ID   uint `gorm:"primaryKey"`
	Name strArr
}

// Value 放到数据库中
func (arr strArr) Value() (driver.Value, error) {
	return strings.Join(arr, ","), nil
}

// Scan 从数据库中读取到结构体中
func (arr *strArr) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	*arr = strings.Split(string(bytes), ",")
	return nil
}

func RunDefinition(db *gorm.DB) {
	db.AutoMigrate(&UserD{})

	//db.Create(&UserD{Name: strArr{"Tom", "Jerry", "Mike"}})

	var user UserD
	db.First(&user, "id = 1")
	fmt.Println(user)

	//sql.NullString
	//sql.NullByte
}
