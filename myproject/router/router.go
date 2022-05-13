package router

import (
	"github.com/gin-gonic/gin"
	"myproject/controller"
	"myproject/middleware"
)

func CollectRouter(r *gin.Engine) *gin.Engine {
	r.Use(middleware.CORSMiddleware(), middleware.RecoveryMiddle()) //全局使用,进行CROS跨域
	r.POST("/api/auth/register", controller.Register)
	r.POST("/api/auth/login", controller.Login)
	//传入中间件AuthMiddleware()
	r.GET("/api/auth/info", middleware.AuthMiddleware(), controller.Info)
	// 创建路由组
	categoryRoutes := r.Group("/category")
	catecon := controller.NewCategoryController()
	categoryRoutes.POST("", catecon.Create)
	categoryRoutes.PUT("/:id", catecon.Update)
	categoryRoutes.DELETE("/:id", catecon.Delete)
	categoryRoutes.GET("/:id", catecon.Show)

	postRoutes := r.Group("/posts")
	//由于使用AuthMiddleware ，所以我们在post的时候需要先获取api中其中一个用户token，通过Auth中间件的认证。一般我们可以重新创建用户获取其token
	// 在Authorization 中bear token中输入token
	postRoutes.Use(middleware.AuthMiddleware())
	postController := controller.NewPostController()
	postRoutes.POST("", postController.Create)
	postRoutes.PUT("/:id", postController.Update) //替换
	postRoutes.GET("/:id", postController.Show)
	postRoutes.DELETE("/:id", postController.Delete)
	postRoutes.POST("/page/list", postController.PageList)
	return r
}
