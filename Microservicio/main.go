package main

import (
	"microservicio/utils/db"
	"microservicio/router"
	"github.com/gin-gonic/gin"
	"microservicio/utils/cache"

	"fmt"
)

var (
	gin_router *gin.Engine
)

func main() {
	gin_router = gin.Default()
	router.MapUrls(gin_router)
	cache.Init_cache()
	err:=db.Init_db()
defer db.Disconect_db()

if err!=nil {
		fmt.Println("Cannot init db")
		fmt.Println(err)
		return;
	}
	fmt.Println("Starting server")
	gin_router.Run(":8090")
}

 