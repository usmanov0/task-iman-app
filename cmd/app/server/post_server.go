package server

import (
	"fmt"
	gr "google.golang.org/grpc"
	"log"
	"net"
	"os"
	adapter2 "test-project-iman/internal/post-crud-service/adapter"
	app2 "test-project-iman/internal/post-crud-service/app"
	grpc2 "test-project-iman/internal/post-crud-service/delivery/grpc"
	"test-project-iman/pkg/common"
	pb2 "test-project-iman/proto/post_proto/crud_grpc/pb"
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

	crudServiceRepo := adapter2.NewPostCrudRepository(db)
	crudUsecase := app2.NewPostCrudUseCase(crudServiceRepo)

	dataCrudGrpc := grpc2.NewCrudServiceServer(crudUsecase)

	listener, err := net.Listen("tcp", os.Getenv("GRPC_PORT2"))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	fmt.Println("listening op port 8081")
	s := gr.NewServer()
	pb2.RegisterCrudServiceServer(s, dataCrudGrpc)
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
