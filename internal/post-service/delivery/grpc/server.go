package grpc

import (
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	pb "test-project-iman/proto/post_proto/post_grpc/pb"

	"test-project-iman/internal/post-service/app"
)

type PostServer struct {
	postCrud app.PostService
	pb.UnimplementedPostServiceServer
}

func NewPostServiceServer(postCrud app.PostService) pb.PostServiceServer {
	return &PostServer{postCrud: postCrud}
}

func (c *PostServer) GetList(ctx context.Context, req *pb.GetPostsList) (*pb.PostList, error) {
	var page, limit int
	posts, err := c.postCrud.GetAllPosts(page, limit)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to get posts: %v", err)
	}

	var pbPosts []*pb.Post
	for _, post := range posts {
		pbPost := &pb.Post{
			Id:     post.Id,
			UserId: post.UserId,
			Title:  post.Title,
			Body:   post.Body,
		}
		pbPosts = append(pbPosts, pbPost)
	}
	return &pb.PostList{Post: pbPosts}, nil
}

func (c *PostServer) GetPost(ctx context.Context, req *pb.PostRequestId) (*pb.Post, error) {
	postId := int(req.Id)

	post, err := c.postCrud.GetOne(postId)

	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Failed to get post: %v", err)
	}

	return &pb.Post{
		Id:     post.Id,
		UserId: post.UserId,
		Title:  post.Title,
		Body:   post.Body,
	}, nil
}

func (c *PostServer) Update(ctx context.Context, req *pb.PostUpdate) (*pb.Empty, error) {
	postId := int(req.Id)

	err := c.postCrud.Update(postId, req.Title, req.Body)
	if err != nil {
		return nil, fmt.Errorf("Failed to update post: %v", err)
	}

	return &pb.Empty{}, nil
}

func (c *PostServer) Delete(ctx context.Context, req *pb.PostRequestId) (*pb.Empty, error) {
	postId := int(req.Id)

	err := c.postCrud.Delete(postId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to delete post: %v", err)
	}
	return &pb.Empty{}, nil
}
