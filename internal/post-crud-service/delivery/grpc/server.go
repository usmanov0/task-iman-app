package grpc

import (
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	pb2 "test-project-iman/proto/post_proto/crud_grpc/pb"

	"test-project-iman/internal/post-crud-service/app"
)

type CrudServer struct {
	postCrud app.PostCrudService
	pb2.UnimplementedCrudServiceServer
}

func NewCrudServiceServer(postCrud app.PostCrudService) pb2.CrudServiceServer {
	return &CrudServer{postCrud: postCrud}
}

func (c *CrudServer) GetList(ctx context.Context, req *pb2.PostRequestPage) (*pb2.PostList, error) {
	if ctx.Err() != nil {
		return nil, status.Errorf(codes.Canceled, "Request canceled")
	}

	posts, err := c.postCrud.GetAllPosts()

	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to get posts: %v", err)
	}

	var pbPosts []*pb2.Post
	for _, post := range posts {
		pbPost := &pb2.Post{
			Id:     post.Id,
			UserId: post.UserId,
			Title:  post.Title,
			Body:   post.Body,
		}
		pbPosts = append(pbPosts, pbPost)
	}
	return &pb2.PostList{Post: pbPosts}, nil
}

func (c *CrudServer) GetPost(ctx context.Context, req *pb2.PostRequestId) (*pb2.Post, error) {
	if ctx.Err() != nil {
		return nil, status.Errorf(codes.Canceled, "Request canceled")
	}

	postId := int(req.Id)

	post, err := c.postCrud.GetOne(postId)

	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Failed to get post: %v", err)
	}

	return &pb2.Post{
		Id:     post.Id,
		UserId: post.UserId,
		Title:  post.Title,
		Body:   post.Body,
	}, nil
}

func (c *CrudServer) Update(ctx context.Context, req *pb2.PostUpdate) (*pb2.Result, error) {
	if ctx.Err() != nil {
		return nil, status.Errorf(codes.Canceled, "Request canceled")
	}

	postId := int(req.Id)

	updateRes, err := c.postCrud.Update(postId, req.Title, req.Body)
	if err != nil {
		return &pb2.Result{
			Success: false,
			Message: fmt.Sprintf("Failed to update post: %v", err),
		}, status.Errorf(codes.Internal, "Failed to update post: %v", err)
	}

	return &pb2.Result{
		Success: true,
		Message: updateRes.Message,
	}, nil
}

func (c *CrudServer) Delete(ctx context.Context, req *pb2.PostRequestId) (*pb2.Empty, error) {
	if ctx.Err() != nil {
		return nil, status.Errorf(codes.Canceled, "Request canceled")
	}

	postId := int(req.Id)

	err := c.postCrud.Delete(postId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to delete post: %v", err)
	}
	return &pb2.Empty{}, nil
}
