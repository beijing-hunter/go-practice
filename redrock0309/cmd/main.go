package main

import (
	"github.com/gin-gonic/gin"
	"redrock0309/router"
)

func main(){

	app := gin.Default()

	router.SetupRouter(app)

	app.Run(":8089")
}
