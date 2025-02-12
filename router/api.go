package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kevinanielsen/go-fast-cdn/handlers"
	dbHandlers "github.com/kevinanielsen/go-fast-cdn/handlers/db"
	dHandlers "github.com/kevinanielsen/go-fast-cdn/handlers/docs"
	iHandlers "github.com/kevinanielsen/go-fast-cdn/handlers/image"
	"github.com/kevinanielsen/go-fast-cdn/util"
)

func AddApiRoutes(r *gin.Engine) {
	api := r.Group("/api")
	api.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "pong")
	})

	cdn := api.Group("/cdn")

	{
		cdn.GET("/size", handlers.GetSizeHandler)
		cdn.GET("/doc/all", dHandlers.HandleAllDocs)
		cdn.GET("/image/all", iHandlers.HandleAllImages)
		cdn.POST("/drop/database", dbHandlers.HandleDropDB)
		cdn.Static("/download/images", util.ExPath+"/uploads/images")
		cdn.Static("/download/docs", util.ExPath+"/uploads/docs")
	}

	upload := cdn.Group("/upload")
	{
		upload.POST("/image", iHandlers.HandleImageUpload)
		upload.POST("/doc", dHandlers.HandleDocUpload)
	}

	delete := cdn.Group("/delete")
	{
		delete.DELETE("/image/:filename", iHandlers.HandleImageDelete)
		delete.DELETE("/doc/:filename", dHandlers.HandleDocDelete)
	}

	rename := cdn.Group("/rename")
	{
		rename.PUT("/image", iHandlers.HandleImageRename)
		rename.PUT("/doc", dHandlers.HandleDocsRename)
	}
}
