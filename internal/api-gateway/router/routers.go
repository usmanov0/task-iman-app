package router

import (
	"github.com/gin-gonic/gin"
	"test-project-iman/internal/api-gateway/delivery/grpc_server_client"
	fetcher "test-project-iman/internal/api-gateway/http"
	post "test-project-iman/internal/api-gateway/http"
)

func SetupRouter(s grpc_server_client.ServiceManager) *gin.Engine {
	router := gin.Default()

	fetchCtrl := fetcher.New(s)

	postCtlr := post.NewPostController(s)

	router.GET("/fetch-data", fetchCtrl.FetchData)

	postRoutes := router.Group("/posts")
	{
		postRoutes.GET("", postCtlr.GetList)
		postRoutes.GET("/id", postCtlr.GetOne)
		postRoutes.PUT("/id", postCtlr.Update)
		postRoutes.DELETE("/id", postCtlr.Delete)
	}

	return router
}
