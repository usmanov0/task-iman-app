package grpc_server_client

import (
	"google.golang.org/grpc"
	"log"
	fetcher "test-project-iman/proto/fetcher_proto/fetcher_grpc/pb"
	postservice "test-project-iman/proto/post_proto/crud_grpc/pb"
)

const (
	service1Address = "post-collector:8080"
	service2Address = "post-service:8081"
)

type ServiceManager interface {
	FetchData() fetcher.CollectorServiceClient
	PostService() postservice.CrudServiceClient
}

type grpcClient struct {
	fetchData  fetcher.CollectorServiceClient
	postClient postservice.CrudServiceClient
}

func (c *grpcClient) FetchData() fetcher.CollectorServiceClient {
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
		fetchData:  fetcher.NewCollectorServiceClient(fetchConn),
		postClient: postservice.NewCrudServiceClient(postClient),
	}
	return s
}
