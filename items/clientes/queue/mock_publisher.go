package queue

import (
	"context"
	"items/dtos"
	"microservicio/dtos"
)

type PublisherMock struct{}

func (PublisherMock) Publish(ctx context.Context, item dtos.ItemDto) error {
	panic("implement me")
}
