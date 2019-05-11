package main

import (
	"./database"
	"./routers"
	"github.com/gin-gonic/gin"
)
func main() {
	database.InitMysql()
	router := routers.InitRouter()
	gin.SetMode(gin.ReleaseMode)
	//静态资源
	router.Static("/static", "./static")
	router.Run(":8081")
}
