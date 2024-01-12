package grpc

import (
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"test-project-iman/internal/post-collector-service/app"
	"test-project-iman/proto/fetcher_proto/fetcher_grpc/pb"
)

type Server struct {
	postService app.PostService
	pb.UnimplementedCollectorServiceServer
}

func NewDataCollectorServer(postService app.PostService) pb.CollectorServiceServer {
	return &Server{postService: postService}
}

func (s *Server) CollectorPosts(ctx context.Context, empty *pb.Empty) (*pb.Result, error) {
	if ctx.Err() != nil {
		return nil, status.Errorf(codes.Canceled, "Request canceled")
	}

	err := s.postService.CollectPosts()

	if err != nil {
		return nil, fmt.Errorf("failed to collect posts: %v", err)
	}

	response := &pb.Result{
		StatusMessage: "Posts Collected successfully and saved to database",
	}

	return response, nil
}
