package router

import (
	ctr "mweibo/controller"
	ctrget "mweibo/controller/get"
	ctrpost "mweibo/controller/post"

	"github.com/gin-gonic/gin"
)

func registerApi(e *gin.Engine) {
	e.GET("/", ctrget.Home)
	//e.GET("/captcha", GetCaptcha)

	e.GET("/register", ctr.RegisterGet)
	e.POST("/register", ctr.RegisterPost)
	e.GET("/login", ctrget.LoginGET)
	e.POST("/login", ctrpost.LoginPOST)
	e.GET("/logout", ctrget.LogoutGET)

	//e.GET("weibo/:id", ctr.DisplayWeibo)
	e.GET("/tag/:tag", ctr.ListTags)

	e.GET("/readweibo/:id", ctrget.DisplayWeibo)
	e.GET("/delweibo/:id", ctrpost.DeleteWeibo)

	e.GET("/createtags", ctrpost.CreateTagsGET)
	e.POST("/createtags", ctrpost.CreateTag)

	//e.GET("/user/:id", AuthUser(ctrget.DisplayUser))
	e.GET("/user/:id", ctrget.UserInfo)

	auth := e.Group("/auth")
	// auth.Use(Auth())
	//auth.Use()
	{
		auth.PUT("/user", ctr.UpdateUserAvatar)
		auth.POST("/bindemail", ctr.BindUserEmail)
		auth.POST("/unbindemail", ctr.UnbindUserEmail)

		auth.GET("/weibo", ctrget.CreateWeiboGET)
		auth.POST("/weibo", ctrpost.CreateWeiboPOST)
		auth.GET("/weibo/:id", ctrget.UpdateWeiboGET)
		auth.POST("/weibo/:id", ctrpost.UpdateWeiboPOST)

		auth.GET("/comment", ctrget.CreateCommentGET)
		auth.POST("/comment", ctrpost.CreateCommentPOST)
		auth.GET("/comment/:id", ctrget.UpdateCommentGET)
		auth.POST("/comment/:id", ctrpost.UpdateCommentPOST)

		// auth.POST("/comment", ctr.CreateComment)
		// auth.POST("/comment/:id", ctr.ReadComment)
		// auth.POST("/comments", ctr.ReadAllComments)

		auth.POST("/tag", ctr.CreateTag)
	}
	admin := e.Group("/admin")
	admin.Use(Admin())
	{
		admin.GET("/users", ctr.ListUsers)
		//admin.GET("/weibos", ctr.ListWeibosByTag)

		//admin.DELETE("/weibo/:id", ctr.DeleteWeibo)
		admin.DELETE("/comment/:id", ctr.DeleteComment)
	}
}
