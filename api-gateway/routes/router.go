package routes

import (
	"api-gateway/internal/handler"
	"api-gateway/middleware"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func NewRouter(service ...interface{}) *gin.Engine {
	ginRouter := gin.Default()
	ginRouter.Use(middleware.Cors(), middleware.InitMiddleware(service), middleware.ErrorMiddleware())
	store := cookie.NewStore([]byte("something-very-secret"))
	ginRouter.Use(sessions.Sessions("mysession", store))
	v1 := ginRouter.Group("/api/v1")
	{
		v1.GET("ping", func(context *gin.Context) {
			context.JSON(200, "success")
		})
		User := &handler.UserService{}
		// 用户服务
		// v1.POST("/user/register", handler.UserRegister)
		v1.POST("/user/login", User.Login)
		v1.POST("/user/create", User.Create)

		// 需要登录保护
		// authed := v1.Group("/")
		// authed.Use(middleware.JWT())
		// {
		// 	// 任务模块
		// 	authed.GET("task", handler.GetTaskList)
		// 	authed.POST("task", handler.CreateTask)
		// 	authed.PUT("task", handler.UpdateTask)
		// 	authed.DELETE("task", handler.DeleteTask)
		// }
	}
	return ginRouter
}
