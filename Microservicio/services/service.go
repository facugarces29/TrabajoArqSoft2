package services

import (
	"microservicio/dtos"
	e "microservicio/utils/errors"
)

type Service interface {
	Get(id string) (dtos.ItemDto, e.ApiError)
	Insert(item dtos.ItemDto) (dtos.ItemDto, e.ApiError)
}
