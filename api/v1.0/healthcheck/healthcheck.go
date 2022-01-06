package healthcheck

import (
	"github.com/gin-gonic/gin"
)

func ApplyRoutes(r *gin.RouterGroup) {
	healthcheck := r.Group("/healthcheck")
	{
		healthcheck.GET("/", getHealthcheck)
		healthcheck.GET("", getHealthcheck)
	}
}
