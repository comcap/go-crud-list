package app

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"meeting-room/config"
	"net/http"

	"meeting-room/app/event"

	"meeting-room/docs"
	eventService "meeting-room/service/event"
)

type App struct {
	event    *event.Controller
	basePath string
}

func New(eventService eventService.Service, conf *config.Config) *App {
	return &App{
		event:    event.New(eventService),
		basePath: conf.AppBasePath,
	}
}

func (app *App) RegisterRoute(router *gin.Engine) *App {
	docs.SwaggerInfo.Title = "Touch Tech API"
	docs.SwaggerInfo.Description = "API Spec Demo."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "http://localhost:8080"
	docs.SwaggerInfo.BasePath = app.basePath

	router.GET("/system/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, nil)
		return
	})

	apiRoutes := router.Group(docs.SwaggerInfo.BasePath)
	{
		//apiRoutes.GET("/companies", app.company.List)
		//apiRoutes.POST("/companies", app.company.Create)
		//apiRoutes.GET("/companies/:id", app.company.Read)
		//apiRoutes.PUT("/companies/:id", app.company.Update)
		//apiRoutes.DELETE("/companies/:id", app.company.Delete)
		//
		//apiRoutes.GET("/staffs", app.staff.Update)
		//apiRoutes.POST("/staffs", app.staff.Create)
		//apiRoutes.GET("/staffs/:id", app.staff.Read)
		//apiRoutes.PUT("/staffs/:id", app.staff.Update)
		//apiRoutes.DELETE("/staffs/:id", app.staff.Delete)

		apiRoutes.GET("/event", app.event.List)
		apiRoutes.POST("/event", app.event.Create)
		apiRoutes.GET("/event/:id", app.event.Read)
		apiRoutes.PUT("/event/:id", app.event.Update)
		apiRoutes.DELETE("/event/:id", app.event.Delete)
	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return app
}
