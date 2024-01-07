package httpServer

import (
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"os"
	"test-project-iman/internal/common"
	adapter2 "test-project-iman/internal/post-collector-service/adapter"
	"test-project-iman/internal/post-collector-service/app"
	"test-project-iman/internal/post-collector-service/post_fetcher_http"
	"test-project-iman/internal/post-crud-service/adapter"
	app2 "test-project-iman/internal/post-crud-service/app"
	"test-project-iman/internal/post-crud-service/post_crud_http"
)

func HttpServer() *chi.Mux {
	db, err := common.ConnectToDb(
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_DATABASE"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
	)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	defer db.Close()

	router := chi.NewRouter()

	postRepo := adapter2.NewPostRepository(db)
	postProvider := adapter2.NewPostCollectorRepository(db)
	postUseCase := app.NewPostService(postRepo, postProvider)
	postHandler := post_fetcher_http.NewPostController(postUseCase)

	postCrudRepo := adapter.NewPostCrudRepository(db)
	postCrudUseCase := app2.NewPostCrudUseCase(postCrudRepo)
	postCrudHandler := post_crud_http.NewPostCrudController(postCrudUseCase)

	router.Get("/", func(w http.ResponseWriter, _ *http.Request) {
		w.Write([]byte("Hello, Chi!"))
	})

	router.Route("/api", func(r chi.Router) {
		r.Route("/posts", func(r chi.Router) {
			r.Get("/get", postHandler.GetPost)
			r.Get("/get-all", postCrudHandler.GetAll)
			r.Get("/get-one", postCrudHandler.GetOne)
			r.Put("/update", postCrudHandler.Update)
			r.Delete("/delete", postCrudHandler.Delete)
		})
	})

	server := &http.Server{Addr: os.Getenv("HTTP_PORT"), Handler: router}
	log.Println("Starting server on port...", os.Getenv("HTTP_PORT"))
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		panic(err)
	}
	defer server.Close()

	return router
}
