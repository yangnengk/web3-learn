package lesson02

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string  `gorm:"size:256"` // 该字段不插入是空字符串
	Email    *string `gorm:"size:100"` // 该字段不插入时是null
	Age      uint    // An unsigned 8-bit integer
	Birthday time.Time
}

// BeforeCreate 在创建记录之前调用 (钩子hook)
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.Name = u.Name + "123"
	return
}

func Run(db *gorm.DB) {
	db.AutoMigrate(&User{})
	// user := User{
	// 	Name:     "Jinzhu",
	// 	Age:      18,
	// 	Birthday: time.Now(),
	// }
	// result := db.Create(&user) //通过数据的指针来创建, 不传指针无法创建记录（报错：panic: reflect: reflect.Value.Set using unaddressable value）
	// fmt.Println(result.RowsAffected)
	// user.ID	// 通过数据的指针来创建, 可以将数据中的ID赋值给user.ID
	// result.Error //返回错误
	// result.RowsAffected //返回受影响的行数
	// result.Statement //返回生成的SQL语句

	// 批量插入
	// users := []*User{
	// 	{Name: "Jinzhu2", Age: 19, Birthday: time.Now()},
	// 	{Name: "Jinzhu3", Age: 20, Birthday: time.Now()},
	// } // 通过数组指针来创建
	// result := db.Create(users)
	// fmt.Println(users[0])
	// fmt.Println(result.RowsAffected)

	// 查询
	// var user User
	// Debug模式会打印出生成的SQL语句
	// result := db.Debug().First(&user) //有排序SELECT * FROM `users` WHERE `users`.`deleted_at` IS NULL ORDER BY `users`.`id` LIMIT 1
	// error := db.Debug().Take(&user).Error //SELECT * FROM `users` WHERE `users`.`deleted_at` IS NULL AND `users`.`id` = 1 LIMIT 1
	// if error != nil {
	// 	panic(error)
	// }

	// 检查 ErrRecordNotFound 错误
	// errors.Is(result.Error, gorm.ErrRecordNotFound)

	// 通过主键查询
	// user.ID = 1
	// db.Debug().First(&user, []int{1, 2, 3}) //SELECT * FROM `users` WHERE `users`.`id` IN (1,2,3) AND `users`.`deleted_at` IS NULL AND `users`.`id` = 1 ORDER BY `users`.`id` LIMIT 1

	// // 查询 全部对象
	// var users []User
	// db.Debug().Find(&users) //SELECT * FROM `users` WHERE `users`.`deleted_at` IS NULL

	var user User
	user.ID = 1
	db.Debug().First(&user, "name = ?", "Jinzhu")
}
