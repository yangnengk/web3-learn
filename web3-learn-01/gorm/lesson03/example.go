package lesson03

import "gorm.io/gorm"

// `User` 属于 `Company`，`CompanyID` 是外键
// type User struct {
// 	gorm.Model
// 	Name      string
// 	CompanyID int // 默认外键
// 	Company   Company
// }
// type User2 struct {
// 	gorm.Model
// 	Name       string
// 	CompanyIDS int
// 	Company    Company `gorm:"foreignKey:CompanyIDS"`
// 	// 使用 CompanyRefer 作为外键
// }

// type User3 struct {
// 	gorm.Model
// 	Name      string
// 	CompanyID string  `gorm:size:32`           // 不指定长度会报错 Error 1170 (42000): BLOB/TEXT column 'company_id' used in key specification without a key length
// 	Company   Company `gorm:"references:Code"` // 使用 Code 作为引用
// }

type Company struct {
	ID uint `gorm:"primaryKey"`
	// Code string `gorm:"unique;size:32"` // 不指定长度会报错 Error 1170 (42000): BLOB/TEXT column 'company_id' used in key specification without a key length
	Name string
}

type Address struct {
	ID      uint `gorm:"primaryKey"`
	Address string
	UserID  uint
}

type Language struct {
	ID   uint `gorm:"primaryKey"`
	Name string
}

type Email struct {
	ID     uint `gorm:"primaryKey"`
	Email  string
	UserID uint
}

type User struct {
	ID             uint `gorm:"primaryKey"`
	Name           string
	BillingAddress Address
	Emails         []Email
	Languages      []Language `gorm:"many2many:user_languages;"` // 多对多关系
	CompanyID      uint       `gorm:unique`
	Company        Company    `gorm:"foreignKey:CompanyID"`
}

func CreateUser(db *gorm.DB) {
	comp := Company{
		Name: "company",
	}
	db.Create(&comp)

	user := User{
		Name:           "lisisisi",
		BillingAddress: Address{Address: "billing address"},
		Emails: []Email{
			{Email: "1234567890@qq.com"},
			{Email: "1234567890@163.com"},
		},
		Languages: []Language{
			{Name: "golang"},
			{Name: "java"},
			{Name: "python"},
		},
		CompanyID: comp.ID,
	}
	db.Create(&user)
}

func Run(db *gorm.DB) {
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Company{}, &Address{}, &Language{}, &Email{})

	// CreateUser(db)

	var user User

	db.First(&user)

	// var langs []Language
	var address Address
	db.Debug().Model(&User{ID: 1}).Association("BillingAddress").Find(&address)
	// db.Debug().Model(&User{ID: 1}).Association("Languages").Append([]Language{{Name: "php"}})

}
