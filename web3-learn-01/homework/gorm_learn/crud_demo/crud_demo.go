package cruddemo

import (
	"fmt"

	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	Name  string `gorm:"type:varchar(20);not null;"`
	Age   int    `gorm:"type:int(3);not null;"`
	Grade string `gorm:"type:varchar(20);not null;"`
}

func Run(db *gorm.DB) {
	db.AutoMigrate(&Student{})
	// 插入一条数据
	// db.Create(&Student{Name: "张三", Age: 20, Grade: "三年级"})
	queryStudent := Student{}
	db.Where("age > ?", 18).Find(&queryStudent)
	fmt.Println(queryStudent)
	db.Model(&Student{}).Where(&Student{Name: "张三"}).Update("grade", "四年级")
	db.Debug().Where("age < ?", 15).Delete(&Student{})
}
