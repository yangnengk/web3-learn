package blog

import (
	"fmt"
	"log"

	"gorm.io/gorm"
)

type User struct {
	ID        uint   `gorm:"primary_key"`
	Name      string `gorm:"type:varchar(20);unique_index"`
	PostCount int    `gorm:"default:0;type:int(10);"`
	Posts     []Post `gorm:"foreignKey:UserID"`
}

type Post struct {
	ID            uint      `gorm:"primary_key"`
	Content       string    `gorm:"type:text"`
	UserID        uint      `gorm:"index;not null"`
	CommentCount  int       `gorm:"default:0;type:int(10);"`
	CommentStatus string    `gorm:type:varchar(20);`
	Comments      []Comment `gorm:"foreignKey:PostID"`
}

type Comment struct {
	ID      uint   `gorm:"primary_key"`
	Content string `gorm:"type:text"`
	PostID  uint   `gorm:"index;not null"`
}

func Run(db *gorm.DB) {
	db.AutoMigrate(&User{}, &Post{}, &Comment{})

	//create
	db.Create(
		&User{
			Name: "张三",
			Posts: []Post{
				{
					Content: "张三的文章1",
					Comments: []Comment{
						{Content: "张三的评论1"},
						{Content: "张三的评论2"},
					},
				},
				{
					Content: "张三的文章2",
					Comments: []Comment{
						{Content: "张三的评论3"},
						{Content: "张三的评论4"},
					},
				},
			},
		},
	)
	db.Create(
		&User{
			Name: "李四",
			Posts: []Post{
				{
					Content: "李四的文章",
					Comments: []Comment{
						{Content: "李四的评论"},
						{Content: "李四的评论2"},
					},
				},
				{
					Content: "李四的文章2",
					Comments: []Comment{
						{Content: "李四的评论3"},
						{Content: "李四的评论4"},
					},
				},
			},
		},
	)

	// query
	var user User
	var posts Post
	var comment Comment
	db.Where("name = ?", "张三").Find(&user)
	fmt.Println(user)
	error := db.Model(&user).Association("Posts").Find(&posts)
	if error != nil {
		panic(error)
	}
	fmt.Println(posts)
	db.Model(&posts).Association("Comments").Find(&comment)
	fmt.Println(comment)

	var user2 []User
	db.Where("name = ?", "李四").Preload("Posts").Preload("Posts.Comments").Find(&user2)
	fmt.Println(user2)

	// 查询 使用Gorm查询评论数量最多的文章信息
	maxSql := `SELECT t.post_id 
		FROM (
			SELECT 
				c.post_id,
				ROW_NUMBER() OVER (ORDER BY c.count DESC) r
			FROM (
				SELECT 
					c.post_id,
					COUNT(*) count
				FROM comments c
				GROUP BY post_id
			) c
		) t WHERE t.r=1`
	db.Raw(maxSql).Scan(&comment)
	fmt.Println(comment)
	db.First(&posts, comment.PostID)
	fmt.Println("max comment post: ", posts)
}

// Hook
func RunHook(db *gorm.DB) {
	db.AutoMigrate(&User{}, &Post{}, &Comment{})

	// 创建文章
	var user User
	db.Where("name = ?", "张三").Find(&user)
	//db.Create(&Post{Content: "张三的可爱文章", UserID: user.ID})
	//db.Create(&Post{Content: "张三的可爱文章2", UserID: user.ID})

	// 删除评论
	// 先查询要删除的评论，然后再删除，确保钩子函数能获取到完整的评论信息
	var comment = Comment{Content: "张三的评论4被删除了"}
	if err := db.First(&comment, 12).Error; err != nil {
		fmt.Println("评论不存在:", err)
		return
	}
	// 更新评论内容
	// 保存更新后的评论，这种方式会触发AfterUpdate钩子
	if err := db.Debug().Model(&comment).Update("content", "张三的评论4被更新了3").Error; err != nil {
		fmt.Println("更新评论失败:", err)
		return
	}

	rows := db.Where(&comment).Find(&comment).RowsAffected
	if rows == 0 {
		log.Println("评论不存在")
		return
	}
	// 删除一定要主键作为条件，不然无法触发钩子函数
	db.Debug().Delete(&comment) // 删除已查询到的评论对象

}

func (p *Post) AfterCreate(tx *gorm.DB) (err error) {
	fmt.Println("创建文章成功")
	tx.Debug().Model(&User{}).Where("id = ?", p.UserID).Update("post_count", gorm.Expr("post_count + ?", 1))
	return
}

func (c *Comment) AfterDelete(tx *gorm.DB) (err error) {
	fmt.Println("删除评论成功，c: ", c)
	var post Post
	res := tx.First(&post, c.PostID)
	if res.RowsAffected == 0 {
		log.Println("文章不存, postID: ", c.PostID)
		return
	}
	if post.CommentCount == 0 {
		tx.Model(&Post{}).Where("id = ?", c.PostID).Update("comment_status", "无评论")
		return
	}
	tx.Model(&Post{}).Where("id = ?", c.PostID).Update("comment_count", gorm.Expr("comment_count - ?", 1))
	return
}

func (c *Comment) AfterUpdate(tx *gorm.DB) error {
	fmt.Println("更新评论成功，c: ", c)
	return nil
}
