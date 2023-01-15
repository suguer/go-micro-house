package routes

import (
	"api-gateway/internal/handler"
	"api-gateway/middleware"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

type Router struct {
	Services map[string]interface{}
}

func NewRouter() *Router {
	return &Router{
		Services: make(map[string]interface{}),
	}
}
func (r *Router) AddService(name string, service interface{}) {
	r.Services[name] = service
}
func (r *Router) Start() *gin.Engine {
	ginRouter := gin.Default()
	ginRouter.Use(middleware.Cors(), middleware.InitMiddleware(r.Services), middleware.ErrorMiddleware())
	store := cookie.NewStore([]byte("something-very-secret"))
	ginRouter.Use(sessions.Sessions("mysession", store))
	v1 := ginRouter.Group("/api/v1")
	{
		v1.GET("ping", func(context *gin.Context) {
			context.JSON(200, "success")
		})
		userGroup := v1.Group("/user")
		{
			User := &handler.UserService{}
			userGroup.POST("/login", User.Login)
			userGroup.POST("/create", User.Create)
			userGroup.POST("/update", User.Update)
		}

		// 用户服务

		houseGroup := v1.Group("/house")
		{
			s := &handler.HouseService{}
			houseGroup.GET("/index", s.Index)
			houseGroupGroup := houseGroup.Group("/group")
			{
				s := &handler.HouseGroupService{}
				houseGroupGroup.GET("/index", s.Index)
				houseGroupGroup.POST("/create", s.Create)

			}
		}
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
