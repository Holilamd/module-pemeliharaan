package routes

import (
	"github.com/Holilamd/module-pemeliharaan/app/handlers"
	"github.com/Holilamd/module-pemeliharaan/app/repositorys"
	"github.com/Holilamd/module-pemeliharaan/app/services"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
)

func RouteUser(db *pgx.Conn, router *gin.Engine) {
	repository := repositorys.NewRepositoryuser(db)
	service := services.NewServiceUser(repository)
	handler := handlers.NewHandlerUser(service)

	api := router.Group("/api/v1/user")
	api.GET("/all", handler.All)
	api.POST("/create", handler.Create)

}
