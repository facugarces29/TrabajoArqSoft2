package services

import (
	"fmt"
	"microservicio/dtos"
	"microservicio/services/repos"
	e "microservicio/utils/errors"
	"net/http"
)

type ImplService struct {
	localCache repos.Repository
	distCache  repos.Repository
	db         repos.Repository
	solr       *repos.SolrClient
}

func NewImplService(
	localCache repos.Repository,
	distCache repos.Repository,
	db repos.Repository,
	solr *repos.SolrClient,
) *ImplService {
	return &ImplService{
		localCache: localCache,
		distCache:  distCache,
		db:         db,
		solr:       solr,
	}
}

func (serv *ImplService) Get(id string) (dtos.ItemDto, e.ApiError) {
	var item dtos.ItemDto
	var apiErr e.ApiError
	var source string

	// try to find it in localCache
	item, apiErr = serv.localCache.Get(id)
	if apiErr != nil {
		if apiErr.Status() != http.StatusNotFound {
			return dtos.ItemDto{}, apiErr
		}
		// try to find it in distCache
		item, apiErr = serv.distCache.Get(id)
		if apiErr != nil {
			if apiErr.Status() != http.StatusNotFound {
				return dtos.ItemDto{}, apiErr
			}
			// try to find it in db
			item, apiErr = serv.db.Get(id)
			if apiErr != nil {
				if apiErr.Status() != http.StatusNotFound {
					return dtos.ItemDto{}, apiErr
				} else {
					fmt.Println(fmt.Sprintf("Not found item %s in any datasource", id))
					apiErr = e.NewNotFoundApiError(fmt.Sprintf("item %s not found", id))
					return dtos.ItemDto{}, apiErr
				}
			} else {
				source = "db"
				defer func() {
					if _, apiErr := serv.distCache.Insert(item); apiErr != nil {
						fmt.Println(fmt.Sprintf("Error trying to save item in distCache %v", apiErr))
					}
					if _, apiErr := serv.localCache.Insert(item); apiErr != nil {
						fmt.Println(fmt.Sprintf("Error trying to save item in localCache %v", apiErr))
					}
				}()
			}
		} else {
			source = "distCache"
			defer func() {
				if _, apiErr := serv.localCache.Insert(item); apiErr != nil {
					fmt.Println(fmt.Sprintf("Error trying to save item in localCache %v", apiErr))
				}
			}()
		}
	} else {
		source = "localCache"
	}

	fmt.Println(fmt.Sprintf("Obtained item from %s!", source))
	return item, nil
}

func (serv *ImplService) Insert(item dtos.ItemDto) (dtos.ItemDto, e.ApiError) {
	result, apiErr := serv.db.Insert(item)
	if apiErr != nil {
		fmt.Println(fmt.Sprintf("Error inserting item in db: %v", apiErr))
		return dtos.ItemDto{}, apiErr
	}
	fmt.Println(fmt.Sprintf("Inserted item in db: %v", result))

	_, apiErr = serv.distCache.Insert(result)
	if apiErr != nil {
		fmt.Println(fmt.Sprintf("Error inserting item in distCache: %v", apiErr))
		return result, nil
	}
	fmt.Println(fmt.Sprintf("Inserted item in distCache: %v", result))

	_, apiErr = serv.localCache.Insert(result)
	if apiErr != nil {
		fmt.Println(fmt.Sprintf("Error inserting item in localCache: %v", apiErr))
		return result, nil
	}
	fmt.Println(fmt.Sprintf("Inserted item in localCache: %v", result))

	apiErr2 := serv.solr.Update()
	if apiErr2 != nil {
		fmt.Println(fmt.Sprintf("Error inserting item in solr: %v", apiErr2))
		return result, nil
	}
	fmt.Println(fmt.Sprintf("Inserted item in solr: %v", result))

	return result, nil
}
