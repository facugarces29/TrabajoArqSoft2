package app

/*
import (
	"fmt"
	"github.com/gin-gonic/gin"
	itemController "microservicio/controllers/item"
)

func MapUrls(router *gin.Engine) {
	// Products Mapping
	router.GET("/items/:id", itemController.Get)
	router.POST("/items", itemController.InsertItem)

	fmt.Println("Finishing mappings configurations")
}
*/

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var (
	router *gin.Engine
)

func init() {

	router = gin.Default()
	router.Use(cors.Default())

}

func StartRoute() {
	MapUrls()
	router.Run(":8080")
}
