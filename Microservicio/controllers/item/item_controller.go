package item

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"microservicio/dtos"
	service "microservicio/services"
	"microservicio/utils/cache"
	"net/http"
)

/*
func Get(c *gin.Context) {

	id:=c.Param("id")

	res:=cache.Get(id)

	if res!="" {
		var bookDtoCache dtos.BookDto
    	json.Unmarshal([]byte(res), &bookDtoCache)
    	fmt.Println("from cache: " + id)
		c.JSON(http.StatusOK,bookDtoCache)
		return


func InsertItem(c *gin.Context) {
	var itemDto dtos.itemDto
	err := c.BindJSON(&itemDto)

	//error parsing json param
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	itemDto, er := service.ItemService.InsertItem(itemDto)

	// Error del insert
	if er != nil {
		c.JSON(er.Status(), er)
		return
	}
	itemDtoStr, _ := json.Marshal(itemDto)
	cache.Set(itemDto.Id, itemDtoStr)
	fmt.Println("save cache: " + itemDto.Id)
	c.JSON(http.StatusCreated, itemDto)
}
