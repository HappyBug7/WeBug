package router

import (
	"back_end/middleware"

	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	r.Use(middleware.Error)

	apiRouter := r.Group("/api")
	{
		apiRouter.POST("/shortcut_add", ctr.Shortcut.Link_add)
		apiRouter.GET("/shortcuts", ctr.Shortcut.Links_get)
		apiRouter.POST("/weather_get", ctr.Weather.Weather_get)
		apiRouter.GET("/weather", ctr.Weather.Weather_broadcast)
		apiRouter.POST("/class_add", ctr.Schedule.Class_add)
		apiRouter.GET("/class", ctr.Schedule.Class_broadcast)
		apiRouter.GET("/classes", ctr.Schedule.Class_get)
		apiRouter.GET("/search", ctr.Search.Suggestion_get)
	}

}
