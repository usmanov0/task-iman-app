package server

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	http2 "net/http"
	"os"
	"test-project-iman/internal/api-gateway/app"
	client "test-project-iman/internal/api-gateway/delivery/grpc_server_client"
	controller "test-project-iman/internal/api-gateway/http"
)

func SetupApiGatewayRouter() http2.Handler {
	client := client.NewApiClient()
	useCase := app.NewUsecase(client.FetchData(), client.PostService())
	controller := controller.NewController(useCase)

	router := chi.NewRouter()

	router.Use(middleware.Logger)

	router.Route("/posts", func(r chi.Router) {
		router.Get("/fetch-posts", controller.CollectPostsHandler)
		router.Get("/get-posts", controller.GetPosts)
		router.Get("/get-post", controller.GetPost)
		router.Put("/update-post", controller.Update)
		router.Delete("/delete-posts", controller.Delete)
	})

	server := &http2.Server{Addr: os.Getenv("HTTP_PORT"), Handler: router}
	log.Println("Starting HTTP server on port...", os.Getenv("HTTP_PORT"))
	if err := server.ListenAndServe(); err != http2.ErrServerClosed {
		panic(err)
	}
	defer server.Close()

	return router
}
