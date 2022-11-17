package services

import (
	"microservicio/dtos"
	e "microservicio/utils/errors"
)

type Service interface {
	GetItemById(id string) (dtos.ItemDto, e.ApiError)
	InsertItem(item dtos.ItemDto) (dtos.ItemDto, e.ApiError)
}
