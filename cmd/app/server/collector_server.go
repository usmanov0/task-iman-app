package server

import (
	"fmt"
	gr "google.golang.org/grpc"
	"log"
	"net"
	"os"
	"test-project-iman/internal/post-collector-service/adapter"
	"test-project-iman/internal/post-collector-service/app"
	"test-project-iman/internal/post-collector-service/delivery/grpc"
	"test-project-iman/pkg/common"
	"test-project-iman/proto/fetcher_proto/fetcher_grpc/pb"
)

func RunGrpcCollectorServer() {
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

	dataFetcherGrpc := grpc.NewDataCollectorServer(usecase)

	listener, err := net.Listen("tcp", os.Getenv("GRPC_PORT1"))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	fmt.Println("listening op port 8080")
	s := gr.NewServer()
	pb.RegisterCollectorServiceServer(s, dataFetcherGrpc)
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
