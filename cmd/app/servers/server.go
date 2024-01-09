package servers

import (
	"fmt"
	origingr "google.golang.org/grpc"
	"log"
	"net"
	"os"
	"test-project-iman/internal/post-collector-service/adapter"
	"test-project-iman/internal/post-collector-service/app"
	"test-project-iman/internal/post-collector-service/delivery/grpc"
	"test-project-iman/internal/post-collector-service/delivery/grpc/fetcher_grpc/pb"
	adapter2 "test-project-iman/internal/post-crud-service/adapter"
	app2 "test-project-iman/internal/post-crud-service/app"
	grpc2 "test-project-iman/internal/post-crud-service/delivery/grpc"
	pb2 "test-project-iman/internal/post-crud-service/delivery/grpc/crud_grpc/pb"
	"test-project-iman/pkg/common"
)

func RunGrpcServer() {
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

	repo := adapter.NewPostRepository(db)
	providerRepo := adapter.NewPostCollectorRepository(db)
	usecase := app.NewPostService(repo, providerRepo)

	crudServiceRepo := adapter2.NewPostCrudRepository(db)
	crudUsecase := app2.NewPostCrudUseCase(crudServiceRepo)

	dataFetcherGrpc := grpc.NewDataCollectorServer(usecase)
	dataCrudGrpc := grpc2.NewCrudServiceServer(crudUsecase)

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	fmt.Println("listening op port 50051")
	s := origingr.NewServer()
	pb.RegisterCollectorServiceServer(s, dataFetcherGrpc)
	pb2.RegisterCrudServiceServer(s, dataCrudGrpc)
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
