package grpc

import (
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"test-project-iman/internal/post-crud-service/app"
	"test-project-iman/internal/post-crud-service/delivery/grpc/crud_grpc/pb"
)

type CrudServer struct {
	postCrud app.PostCrudService
	pb.UnimplementedCrudServiceServer
}

func NewCrudServiceServer(postCrud app.PostCrudService) pb.CrudServiceServer {
	return &CrudServer{postCrud: postCrud}
}

func (c *CrudServer) GetList(ctx context.Context, _ *pb.Empty) (*pb.PostList, error) {
	if ctx.Err() != nil {
		return nil, status.Errorf(codes.Canceled, "Request canceled")
	}

	posts, err := c.postCrud.GetAllPosts()

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

func (c *CrudServer) GetPost(ctx context.Context, req *pb.PostRequestId) (*pb.Post, error) {
	if ctx.Err() != nil {
		return nil, status.Errorf(codes.Canceled, "Request canceled")
	}

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

func (c *CrudServer) Update(ctx context.Context, req *pb.PostUpdate) (*pb.Result, error) {
	if ctx.Err() != nil {
		return nil, status.Errorf(codes.Canceled, "Request canceled")
	}

	postId := int(req.Id)

	updateRes, err := c.postCrud.Update(postId, req.Title, req.Body)
	if err != nil {
		return &pb.Result{
			Success: false,
			Message: fmt.Sprintf("Failed to update post: %v", err),
		}, status.Errorf(codes.Internal, "Failed to update post: %v", err)
	}

	return &pb.Result{
		Success: true,
		Message: updateRes.Message,
	}, nil
}

func (c *CrudServer) Delete(ctx context.Context, req *pb.PostRequestId) (*pb.Empty, error) {
	if ctx.Err() != nil {
		return nil, status.Errorf(codes.Canceled, "Request canceled")
	}

	postId := int(req.Id)

	err := c.postCrud.Delete(postId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to delete post: %v", err)
	}
	return &pb.Empty{}, nil
}
