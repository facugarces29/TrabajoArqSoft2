package services

import (
	itemDao "microservicio/daos/item"
	"microservicio/dtos"
	model "microservicio/models"
	e "microservicio/utils/errors"
)

type itemService struct{}

type itemServiceInterface interface {
	//GetBook(id string) (dtos.BookDto, e.ApiError)

	InsertItem(itemDto dtos.itemDto) (dtos.itemDto, e.ApiError)
}

var (
	itemService itemServiceInterface
)

func init() {
	itemService = &itemService{}
}

func (s *itemService) InsertItem(itemDto dtos.itemDto) (dtos.itemDto, e.ApiError) {

	item model.item 
	var itemDto dtos.itemDto
	item.Tittle = itemDto.Tittle
	item.Seller = itemDto.Seller
	item.Price = itemDto.Price
	item.Currency = itemDto.Currency
	item.Pictures = itemDto.Pictures
	item.Description = itemDto.Description
	item.State = itemDto.State
	item.City = itemDto.City
	item.Street = itemDto.Street
	item.Number = itemDto.Number

	item = itemDao.InsertItem(item)
	if item.Id.Hex() == "000000000000000000000000" {
		return itenmDto, e.NewBadRequestApiError("error insert item not found")
	}
	
	itemDto.Id = item.Id.Hex()
	
	return itemDto, nil
}

/*func (s *bookService) InsertBook(bookDto dtos.BookDto) (dtos.BookDto, e.ApiError) {

	var book model.Book

	book.Name = bookDto.Name

	book = bookDao.Insert(book)

	if book.Id.Hex() == "000000000000000000000000" {
		return bookDto, e.NewBadRequestApiError("error in insert")
	}
	bookDto.Id = book.Id.Hex()

	return bookDto, nil
}
*/