package repos

import (
	"context"
	"fmt"
	"microservicio/dtos"
	e "microservicio/utils/errors"
	"time"

	"github.com/karlseguin/ccache/v2"
)

type RepositoryCCache struct {
	Client     *ccache.Cache
	DefaultTTL time.Duration
}

func NewCCache(maxSize int64, itemsToPrune uint32, defaultTTL time.Duration) *RepositoryCCache {
	client := ccache.New(ccache.Configure().MaxSize(maxSize).ItemsToPrune(itemsToPrune))
	fmt.Println("[CCache] Initialized")
	return &RepositoryCCache{
		Client:     client,
		DefaultTTL: defaultTTL,
	}
}

func (repo *RepositoryCCache) GetItemById(ctx context.Context, id string) (dtos.ItemDto, e.ApiError) {
	item := repo.Client.Get(id)
	if item == nil {
		return dtos.ItemDto{}, e.NewNotFoundApiError(fmt.Sprintf("item %s not found", id))
	}
	if item.Expired() {
		return dtos.ItemDto{}, e.NewNotFoundApiError(fmt.Sprintf("item %s not found", id))
	}
	return item.Value().(dtos.ItemDto), nil
}

func (repo *RepositoryCCache) InsertItem(ctx context.Context, item dtos.ItemDto) (dtos.ItemDto, e.ApiError) {
	repo.Client.Set(item.Id, item, repo.DefaultTTL)
	return item, nil
}
