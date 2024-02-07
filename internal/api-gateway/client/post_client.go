package client

import (
	"google.golang.org/grpc"
	"log"
	collect "test-project-iman/proto/collector_proto/collector_grpc/pb"
	postservice "test-project-iman/proto/post_proto/post_grpc/pb"
)

const (
	service1Address = "post-collector:8080"
	service2Address = "post-service:8081"
)

type ServiceManager interface {
	FetchData() collect.CollectorServiceClient
	PostService() postservice.PostServiceClient
}

type grpcClient struct {
	fetchData  collect.CollectorServiceClient
	postClient postservice.PostServiceClient
}

func (c *grpcClient) FetchData() collect.CollectorServiceClient {
	return c.fetchData
}

func (c *grpcClient) PostService() postservice.PostServiceClient {
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
		fetchData:  collect.NewCollectorServiceClient(fetchConn),
		postClient: postservice.NewPostServiceClient(postClient),
	}
	return s
}
