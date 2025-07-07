package lesson01

import (
	"database/sql"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID           uint           // Standard field for the primary key
	Name         string         `gorm:"size:256"` // 该字段不插入是空字符串
	Email        *string        `gorm:"size:100"` // 该字段不插入时是null
	Age          uint           // An unsigned 8-bit integer
	Birthday     *time.Time     // A pointer to time.Time, can be null
	MemberNumber sql.NullString // Uses sql.NullString to handle nullable strings(和 *string的区别是，指定Valid=true 如果这个字段是空，那么不会插入null)
	ActivateAt   sql.NullTime   // Uses sql.NullTime for nullable time fields
	CreatedAt    time.Time      // Automatically managed by GORM for creation time
	UpdatedAt    time.Time      // Automatically managed by GORM for update time
	ignore       string         // fields that aren't exported are ignored(因为不是大写开头)

}

type Member struct {
	gorm.Model
	Name string
	Age  uint8
}

type Author struct {
	Name  string
	Email string
}

type Blog2 struct {
	ID      int64
	Author  Author `gorm:"embedded;embeddedPrefix:author_"` // 嵌入字段,embeddedPrefix:author_，那么字段名会变成author_name
	Upvotes int32  `gorm:"column:votes"`                    // 重命名列
}

func Run(db *gorm.DB) {
	// db.AutoMigrate(&User{})

	// user := &User{}
	// user.MemberNumber.Valid = true // 如果这个字段是空，那么也会插入空字符
	// // user.ActivateAt.Valid = true	//sql.NullTime valid = true 时，如果这个字段是空，会报错 Incorrect datetime value: '0000-00-00' for column 'activate_at'
	// db.Create(user)

	// db.AutoMigrate(&Member{})
	// member := &Member{}
	// db.Create(member)
	db.AutoMigrate(&Blog2{})
}
