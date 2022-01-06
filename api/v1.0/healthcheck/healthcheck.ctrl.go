package healthcheck

import "github.com/gin-gonic/gin"

func getHealthcheck(c *gin.Context) {
	c.JSON(200, map[string]string{
		"status": "ok",
	})
}
