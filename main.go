package main

import (
	"github.com/eahrend/go-gin-crud/api"
	mw "github.com/eahrend/go-gin-crud/middleware"
	"github.com/gin-gonic/gin"
	"github.com/volatiletech/sqlboiler/boil"
)

func main() {
	app := gin.Default()
	boil.SetDB(pool)
	app.Use(mw.DB(pool))
	api.ApplyRoutes(app)
	panic(app.Run(":" + port))
}
