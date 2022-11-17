package app

import (
	controllers "items/controllers"
	service "items/services"
	"items/services/repos"
	"time"
)

type Dependecias struct {
	ItemController *controllers.controller
}

func BuildDependecias() *Dependecias {
	// Repo
	ccache := repos.NewCCache(1000, 100, 30*time.Second)
	memcached := repos.NewMemcached("localhost", 11211)
	mongo := repos.NewMongoDB("localhost", 27017, "avisos")
	solr := repos.NewSolrClient("localhost", 8983, "avisos")

	// Services
	service := service.NewImplService(ccache, memcached, mongo, solr)

	// Controllers
	controller := controllers.NewController(service)

	return &Dependecias{
		ItemController: controller,
	}
}
