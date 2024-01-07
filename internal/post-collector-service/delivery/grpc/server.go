package grpc

import (
	"context"
	"fmt"
	"test-project-iman/internal/post-collector-service/app"
	"test-project-iman/internal/post-collector-service/delivery/grpc/fetcher_grpc/pb"
)

type Server struct {
	postService app.PostService
	pb.UnimplementedCollectorServiceServer
}

func NewDataCollectorServer(postService app.PostService) *Server {
	return &Server{postService: postService}
}

func (s *Server) CollectPosts(ctx context.Context, _ *pb.Empty) (*pb.Result, error) {
	err := s.postService.CollectPosts()

	if err != nil {
		return &pb.Result{}, fmt.Errorf("failed to collect posts: %v", err)
	}

	response := &pb.Result{
		StatusMessage: "Posts collected successfully",
	}

	return response, nil
}
