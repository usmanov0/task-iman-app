package app

import (
	"context"
	"test-project-iman/internal/api-gateway/model"
	fetcher "test-project-iman/proto/fetcher_proto/fetcher_grpc/pb"
	postservice "test-project-iman/proto/post_proto/crud_grpc/pb"
)

type Service interface {
	FetcherService
	PostService
}

type FetcherService interface {
	FetchPosts() error
}

type PostService interface {
	GetList() ([]model.Posts, error)
	GetById(id int) (*model.Post, error)
	Update(postId int, title, body string) (*model.PostUpdateResponse, error)
	Delete(postId int) error
}

type service struct {
	fetcherClient fetcher.CollectorServiceClient
	postClient    postservice.CrudServiceClient
}

func NewUsecase(
	fClient fetcher.CollectorServiceClient,
	mClient postservice.CrudServiceClient,
) Service {
	return &service{
		fetcherClient: fClient,
		postClient:    mClient,
	}
}

func (s *service) FetchPosts() error {
	ctx := context.Background()

	req := fetcher.Empty{}

	_, err := s.fetcherClient.CollectorPosts(ctx, &req)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) GetList() ([]model.Posts, error) {
	ctx := context.Background()
	req := postservice.PostRequestPage{}

	res, err := s.postClient.GetList(ctx, &req)

	if err != nil {
		return nil, err
	}

	var posts []model.Posts
	for _, p := range res.Post {
		mappedPost := parseToModel(int(p.Id), int(p.UserId), p.Title, p.Body)
		posts = append(posts, model.Posts{*mappedPost})
	}

	return posts, nil
}

func (s *service) GetById(id int) (*model.Post, error) {
	ctx := context.Background()
	req := postservice.PostRequestId{
		Id: int64(id),
	}

	res, err := s.postClient.GetPost(ctx, &req)

	if err != nil {
		return nil, err
	}
	post := parseToModel(int(res.Id), int(res.UserId), res.Title, res.Body)

	return post, nil
}

func (s *service) Update(postId int, title, body string) (*model.PostUpdateResponse, error) {
	ctx := context.Background()
	req := postservice.PostUpdate{
		Id:    int64(postId),
		Title: title,
		Body:  body,
	}

	res, err := s.postClient.Update(ctx, &req)

	if err != nil {
		return nil, err
	}
	updatePost := model.PostUpdateResponse{res.Success, res.Message}

	return &updatePost, nil
}

func (s *service) Delete(postId int) error {
	ctx := context.Background()
	req := postservice.PostRequestId{
		Id: int64(postId),
	}

	_, err := s.postClient.Delete(ctx, &req)
	if err != nil {
		return err
	}
	return nil
}

func parseToModel(id, userId int, title, body string) *model.Post {
	return &model.Post{
		ID:     id,
		UserId: userId,
		Title:  title,
		Body:   body,
	}
}
