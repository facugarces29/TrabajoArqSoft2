package repos

import (
	"microservicio/dtos"
	"microservicio/utils/errors"
)

type Repository interface {
	Get(id string) (dtos.ItemDto, errors.ApiError)
	Insert(item dtos.ItemDto) (dtos.ItemDto, errors.ApiError)
	Update(item dtos.ItemDto) (dtos.ItemDto, errors.ApiError)
	Delete(id string) errors.ApiError
}
