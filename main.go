package main

import (
	"github.com/Holilamd/module-pemeliharaan/config"
	"github.com/Holilamd/module-pemeliharaan/routes"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {
	db := config.Connect()
	router := gin.Default()
	routes.RouteUser(db, router)
	err := router.Run()
	if err != nil {
		defer logrus.Error("Server is not running")
		logrus.Fatal(err)
	}
}
