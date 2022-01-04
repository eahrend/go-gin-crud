package entries

import (
	"github.com/gin-gonic/gin"
)

func ApplyRoutes(r *gin.RouterGroup) {
	entries := r.Group("/entries")
	{

		entries.GET("/:id", getEntry)
		entries.POST("/:value", postEntry)
	}
}
