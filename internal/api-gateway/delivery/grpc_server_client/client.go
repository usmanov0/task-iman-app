package grpc_server_client

import (
	"google.golang.org/grpc"
	"log"
	"test-project-iman/internal/api-gateway/delivery/grpc_server_client/fetch_datas/pb"
	postservice "test-project-iman/internal/api-gateway/delivery/grpc_server_client/post_service/pb"
)

const (
	service1Address = "post-collector:8080"
	service2Address = "post_service:8081"
)

type ServiceManager interface {
	FetchData() pb.CollectorServiceClient
	PostService() postservice.CrudServiceClient
}

type grpcClient struct {
	fetchData  pb.CollectorServiceClient
	postClient postservice.CrudServiceClient
}

func (c *grpcClient) FetchData() pb.CollectorServiceClient {
	return c.fetchData
}

func (c *grpcClient) PostService() postservice.CrudServiceClient {
	return c.postClient
}

func NewApiClient() ServiceManager {
	fetchConn, err := grpc.Dial(service1Address,
		grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed connect to gRPC service #1: %v", err)
	}

	postClient, err := grpc.Dial(service2Address,
		grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Failed connect to gRPC service #2: %v", err)
	}

	s := &grpcClient{
		fetchData:  pb.NewCollectorServiceClient(fetchConn),
		postClient: postservice.NewCrudServiceClient(postClient),
	}
	return s
}
