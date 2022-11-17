package app

import (
	"github.com/gin-gonic/gin"
)

func StartRoute() {
	router := gin.Default()
	deps := BuildDependecias()
	MapUrls(router, deps)
	_ = router.Run(":8090")
}
