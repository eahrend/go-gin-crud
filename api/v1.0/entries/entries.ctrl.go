package entries

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/eahrend/go-gin-crud/models"
	"github.com/gin-gonic/gin"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func getEntry(c *gin.Context) {
	db := c.MustGet("boiler").(boil.ContextExecutor)
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{
			"error":    fmt.Sprintf("unable to get entry with id %s", id),
			"returned": err.Error(),
		})
		return
	}
	entry, err := models.FindEntry(c, db, idInt)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{
			"error":    fmt.Sprintf("unable to get entry with id %s", id),
			"returned": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, map[string]string{
		"value": entry.Value.String,
	})
}

func postEntry(c *gin.Context) {
	db := c.MustGet("boiler").(boil.ContextExecutor)
	val := c.Param("value")
	newEntry := models.Entry{
		Value: null.String{
			String: val,
			Valid:  true,
		},
	}
	err := newEntry.Insert(c, db, boil.Infer())
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{
			"error":    "unable to insert value",
			"returned": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, map[string]int{
		"id": newEntry.ID,
	})
}
