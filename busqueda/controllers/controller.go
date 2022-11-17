package controllers

import (
	"microservicio/dtos"
	service "microservicio/services"
	e "microservicio/utils/errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type controller struct {
	service service.Service
}

func NewController(service service.Service) *controller {
	return &controller{
		service: service,
	}
}

func (Contr *controller) Get(c *gin.Context) {
	item, apiErr := Contr.service.Get(c.Param("id"))
	if apiErr != nil {
		c.JSON(apiErr.Status(), apiErr)
		return
	}
	c.JSON(http.StatusOK, item)
}

func (contr *controller) Insert(c *gin.Context) {
	var item dtos.ItemDto
	if err := c.BindJSON(&item); err != nil {
		apiErr := e.NewBadRequestApiError(err.Error())
		c.JSON(apiErr.Status(), apiErr)
		return
	}

	item, apiErr := contr.service.Insert(item)
	if apiErr != nil {
		c.JSON(apiErr.Status(), apiErr)
		return
	}

	c.JSON(http.StatusCreated, item)
}
