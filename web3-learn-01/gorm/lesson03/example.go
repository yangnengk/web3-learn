package lesson03

import "gorm.io/gorm"

// `User` 属于 `Company`，`CompanyID` 是外键
type User struct {
	gorm.Model
	Name      string
	CompanyID int // 默认外键
	Company   Company
}
type User2 struct {
	gorm.Model
	Name       string
	CompanyIDS int
	Company    Company `gorm:"foreignKey:CompanyIDS"`
	// 使用 CompanyRefer 作为外键
}

type User3 struct {
	gorm.Model
	Name      string
	CompanyID string  `gorm:size:32`           // 不指定长度会报错 Error 1170 (42000): BLOB/TEXT column 'company_id' used in key specification without a key length
	Company   Company `gorm:"references:Code"` // 使用 Code 作为引用
}

type Company struct {
	ID   int
	Code string `gorm:"unique;size:32"` // 不指定长度会报错 Error 1170 (42000): BLOB/TEXT column 'company_id' used in key specification without a key length
	Name string
}

func Run(db *gorm.DB) {
	db.AutoMigrate(&User3{}, &Company{})
}
