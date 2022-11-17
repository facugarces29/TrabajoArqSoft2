package app

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

//Importa direcciones de ARCHIVOS a obtener

func MapUrls(router *gin.Engine, dependecias *Dependecias) {
	// Mapa de los items
	// Acciones con postMan

	router.GET("/Microservicio/items/:id", dependecias.ItemController.Get)
	router.POST("/Microservicio/items", dependecias.ItemController.Insert)
	fmt.Println("Terminado configuracion de mapeo")
}
