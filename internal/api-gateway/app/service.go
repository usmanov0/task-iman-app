package app

import (
	"context"
	"test-project-iman/internal/api-gateway/model"
	collector "test-project-iman/proto/collector_proto/collector_grpc/pb"
	postservice "test-project-iman/proto/post_proto/post_grpc/pb"
)

type Service interface {
	CollectorService
	PostService
}

type CollectorService interface {
	CollectPosts() error
}

type PostService interface {
	GetList(page, limit int) ([]model.Posts, error)
	GetById(id int) (*model.Post, error)
	Update(postId int, title, body string) error
	Delete(postId int) error
}

type service struct {
	collectorClient collector.CollectorServiceClient
	postClient      postservice.PostServiceClient
}

func NewUsecase(
	fClient collector.CollectorServiceClient,
	mClient postservice.PostServiceClient,
) Service {
	return &service{
		collectorClient: fClient,
		postClient:      mClient,
	}
}

func (s *service) CollectPosts() error {
	ctx := context.Background()

	req := collector.Empty{}

	_, err := s.collectorClient.CollectorPosts(ctx, &req)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) GetList(page, limit int) ([]model.Posts, error) {
	ctx := context.Background()
	req := postservice.GetPostsList{
		Page:  int64(page),
		Limit: int64(limit),
	}

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

func (s *service) Update(postId int, title, body string) error {
	ctx := context.Background()
	req := postservice.PostUpdate{
		Id:    int64(postId),
		Title: title,
		Body:  body,
	}

	_, err := s.postClient.Update(ctx, &req)

	if err != nil {
		return err
	}

	return nil
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
