package queue

import (
	"context"
	"microservicio/dtos"
)

type Publisher interface {
	Publisher(ctx context.Context, item dtos.ItemDto) error
}
