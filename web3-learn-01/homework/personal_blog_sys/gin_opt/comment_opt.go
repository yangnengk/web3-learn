package gin_opt

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/learn/personal_blog_sys/data_model"
	"github.com/learn/personal_blog_sys/data_rm"
)

// 评论操作
func CommentOpt(eg *gin.Engine) {
	commentGroup := eg.Group("/comment")

	{
		commentGroup.POST("/add", func(c *gin.Context) {
			var comment data_model.Comment
			if err := c.ShouldBind(&comment); err != nil {
				log.Println("comment bind error: ", err)
				c.Error(fmt.Errorf("comment bind error: "))
				return
			}
			comment.UserID = c.MustGet("user").(uint)
			data_rm.CreateComment(comment)
			c.JSON(200, gin.H{
				"status":  "success",
				"message": "添加评论成功",
			})
		})

		commentGroup.GET("/list", func(c *gin.Context) {
			comments := data_rm.QueryCommentList()
			c.JSON(200, gin.H{
				"status":  "true",
				"message": "获取评论列表成功",
				"data":    comments,
			})
		})

	}
}
