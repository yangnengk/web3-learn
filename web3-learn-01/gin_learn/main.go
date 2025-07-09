package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type User struct {
	Username string `json:"u" form:"u"`
	Password string `json:"p" form:"p"`
}

func main() {
	gIn := gin.Default()
	//gIn.GET("/", func(context *gin.Context) {
	//	context.String(200, "Hello, World!")
	//})
	//
	//gr1 := gIn.Group("/v1")
	//{
	//	gr1.GET("/user", func(context *gin.Context) {
	//		context.String(200, "v1 user")
	//	})
	//}
	//
	//gr2 := gIn.Group("/v2")
	//{
	//	gr2.GET("/user", func(context *gin.Context) {
	//		context.String(http.StatusOK, "v2 user")
	//	})
	//}
	//
	//// 重定向
	//gIn.GET("/member", func(context *gin.Context) {
	//	context.Redirect(http.StatusFound, "/v1/user")
	//	//context.Redirect(http.StatusMovedPermanently, "www.baidu.com")
	//})
	//// 内部重定向，url不会变化成/v1/user
	//gIn.GET("/member2", func(context *gin.Context) {
	//	context.Request.URL.Path = "/v1/user"
	//	gIn.HandleContext(context)
	//})
	//
	//gIn.GET("/json", func(context *gin.Context) {
	//	context.JSON(http.StatusOK, gin.H{"name": "张三"})
	//})
	//
	//// 静态文件
	//gIn.Static("/static", "./static")
	////gIn.StaticFS("/static", http.Dir("static"))
	//gIn.StaticFile("/fs", "/static/11.txt") // 指定单个文件

	//gIn.GET("/json", func(context *gin.Context) {
	//	user := User{
	//		Username: "张三",
	//		Password: "<PASSWORD>",
	//	}
	//	context.JSON(200, &user)
	//})
	//
	//// html
	//gIn.LoadHTMLGlob("./templates/**/*")
	//gIn.GET("/index", func(context *gin.Context) {
	//	context.HTML(http.StatusOK, "sub1/a.tmpl", gin.H{
	//		"title": "首页",
	//	})
	//})

	gIn.GET("/user/:u/:p", func(context *gin.Context) {
		username := context.Param("u")
		password := context.Param("p")
		context.JSON(http.StatusOK, gin.H{
			"username": username,
			"password": password,
		})
	})

	gIn.GET("/user", func(context *gin.Context) {
		//username := context.Query("u")
		//password := context.Query("p")
		user := &User{}
		err := context.ShouldBind(user)
		if err != nil {
			context.JSON(http.StatusOK, gin.H{
				"msg": err.Error(),
			})
			return
		}
		context.JSON(http.StatusOK, gin.H{
			"username": user.Username,
			"password": user.Password,
		})
	})

	gIn.POST("/userForm", func(context *gin.Context) {
		//username := context.PostForm("u")
		//password := context.PostForm("p")
		user := &User{}
		err := context.ShouldBind(user)
		if err != nil {
			context.JSON(http.StatusOK, gin.H{
				"msg": err.Error(),
			})
			return
		}
		context.JSON(http.StatusOK, gin.H{
			"username": user.Username,
			"password": user.Password,
		})
	})

	// 表单

	err := gIn.Run(":8080") // :8080等于 0.0.0.0:8080
	if err != nil {
		panic(err)
	}

}
