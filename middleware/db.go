package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/volatiletech/sqlboiler/boil"
)

func DB(executor boil.ContextExecutor) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("boiler", executor)
		c.Next()
	}
}
