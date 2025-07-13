package gin_opt

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/learn/personal_blog_sys/data_model"
	"github.com/learn/personal_blog_sys/data_rm"
	"github.com/learn/personal_blog_sys/jwt_util"
	"golang.org/x/crypto/bcrypt"
)

/*
用户操作
*/
func UserOperate(eg *gin.Engine) {

	userGroup := eg.Group("/user")

	{
		// 注册
		userGroup.POST("/registry", func(c *gin.Context) {
			var user data_model.User
			if err := c.ShouldBind(&user); err != nil {
				log.Println("user bind error: ", err)
				c.Error(fmt.Errorf("user bind error: "))
				return
			}
			if user.Email == "" {
				c.Error(fmt.Errorf("email 不能为空"))
				return
			}

			bcryptPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
			if err != nil {
				c.Error(fmt.Errorf("bcrypt error: ", err))
				return
			}
			user.Password = string(bcryptPassword)

			data_rm.CreateUser(user)

			c.JSON(200, gin.H{
				"status":  "success",
				"message": "注册成功",
			})
		})

		// 登录
		userGroup.POST("/login", func(c *gin.Context) {
			var user data_model.User
			if err := c.ShouldBind(&user); err != nil {
				log.Println("user bind error: ", err)
				c.Error(fmt.Errorf("user bind error: ", err))
			}
			storedUser := data_rm.GetUser(user.Username)
			if storedUser.Username == "" {
				c.Error(fmt.Errorf("用户不存在"))
				return
			}

			err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password))
			if err != nil {
				c.Error(fmt.Errorf("密码错误"))
				return
			}

			jwt, err := jwt_util.GenerateJWT(storedUser.ID, storedUser.Username, storedUser.Email)
			if err != nil {
				c.Error(fmt.Errorf("生成token失败"))
				return
			}

			c.JSON(http.StatusOK, gin.H{
				"message": "login success",
				"status":  "success",
				"jwt":     jwt,
			})
		})

		// 查询用户信息
		userGroup.GET("/info", func(c *gin.Context) {
			username := c.Query("username")
			if username == "" {
				c.Error(fmt.Errorf("username 不能为空"))
				return
			}
			userQuery := data_rm.QueryUserInfo(username)
			c.JSON(http.StatusOK, gin.H{
				"message": "success",
				"data":    userQuery,
			})
		})
	}

}
