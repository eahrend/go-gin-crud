package apiv1

import (
	"github.com/eahrend/go-gin-crud/api/v1.0/entries"
	"github.com/eahrend/go-gin-crud/api/v1.0/healthcheck"
	"github.com/gin-gonic/gin"
)

func ApplyRoutes(r *gin.RouterGroup) {
	v1 := r.Group("/v1")
	entries.ApplyRoutes(v1)
	healthcheck.ApplyRoutes(v1)
}
