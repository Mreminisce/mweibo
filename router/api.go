package router

import (
	ctr "mweibo/controller"

	"github.com/gin-gonic/gin"
)

func registerApi(e *gin.Engine) {
	e.GET("/", ctr.Home)
	e.GET("/captcha", GetCaptcha)

	e.GET("/register", ctr.RegisterGet)
	e.POST("/register", ctr.RegisterPost)
	e.GET("/login", ctr.LoginGet)
	e.POST("/login", ctr.LoginPost)
	e.GET("/logout", ctr.Logout)

	e.GET("weibo/:id", ctr.DisplayWeibo)
	e.GET("tag/:tag", ctr.DisplayTag)

	auth := e.Group("/auth")
	auth.Use(Auth())
	{
		auth.PUT("/user", ctr.UpdateUserAvatar)
		auth.POST("/bindemail", ctr.BindUserEmail)
		auth.POST("/unbindemail", ctr.UnbindUserEmail)

		auth.GET("/weibo", ctr.CreateWeiboGet)
		auth.POST("/weibo", ctr.CreateWeiboPost)
		auth.GET("/weibo/:id", ctr.UpdateWeiboGet)
		auth.POST("/weibo/:id", ctr.UpdateWeibo)

		auth.POST("/comment", ctr.CreateComment)
		auth.POST("/comment/:id", ctr.ReadComment)
		auth.POST("/comments", ctr.ReadAllComments)

		auth.POST("/tag", ctr.CreateTag)
	}
	admin := e.Group("/admin")
	admin.Use(Admin())
	{
		admin.GET("/users", ctr.ListUsers)
		admin.GET("/weibos", ctr.ListWeibos)

		admin.DELETE("/weibo/:id", ctr.DeleteWeibo)
		admin.DELETE("/comment/:id", ctr.DeleteComment)
	}
}
