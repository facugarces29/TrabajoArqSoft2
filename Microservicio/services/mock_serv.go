package services

import (
	"microservicio/dtos"
	e "microservicio/utils/errors"
)

type MockServ struct{}

func NewMockServ() MockServ {
	return MockServ{}
}

func (MockServ) Get(id string) (dtos.ItemDto, e.ApiError) {
	return dtos.ItemDto{
		Id:     "2901",
		Tittle: "Mocked item",
	}, nil
}

func (MockServ) Insert(item dtos.ItemDto) (dtos.ItemDto, e.ApiError) {
	return dtos.ItemDto{
		Id:     "2901",
		Tittle: item.Tittle,
	}, nil
}
