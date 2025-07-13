package gin_opt

import (
	"fmt"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/learn/personal_blog_sys/data_model"
	"github.com/learn/personal_blog_sys/data_rm"
)

// 文章操作
func PostOpt(eg *gin.Engine) {
	postGroup := eg.Group("/post")

	{
		// 添加文章
		postGroup.POST("/add", func(c *gin.Context) {
			var post data_model.Post
			if err := c.ShouldBind(&post); err != nil {
				log.Println("post bind error: ", err)
				c.Error(fmt.Errorf("post bind error: "))
				return
			}
			post.UserID = c.MustGet("user").(uint)
			data_rm.CreatePost(post)
			c.JSON(200, gin.H{
				"status":  "success",
				"message": "添加文章成功",
			})
		})

		// 获取文章列表
		postGroup.GET("/list", func(c *gin.Context) {
			posts := data_rm.QueryPostList()
			c.JSON(200, gin.H{
				"status":  "true",
				"message": "获取文章列表成功",
				"data":    posts,
			})
		})

		// 更新文章
		postGroup.POST("/update", func(c *gin.Context) {
			var post data_model.Post
			if err := c.ShouldBind(&post); err != nil {
				log.Println("post bind error: ", err)
				c.Error(fmt.Errorf("post bind error: "))
				return
			}
			post.UserID = c.MustGet("user").(uint)
			data_rm.UpdatePost(post)
			c.JSON(200, gin.H{
				"status":  "success",
				"message": "更新文章成功",
			})
		})

		// 删除文章
		postGroup.DELETE("/delete", func(c *gin.Context) {
			id := c.Query("id")
			var post data_model.Post
			Id, _ := strconv.Atoi(id)
			post.ID = uint(Id)
			post.UserID = c.MustGet("user").(uint)
			data_rm.DeletePost(post)
			c.JSON(200, gin.H{
				"status":  "success",
				"message": "删除文章成功",
			})
		})
	}
}
