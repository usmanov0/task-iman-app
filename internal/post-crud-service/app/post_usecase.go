package app

import (
	"fmt"
	"test-project-iman/internal/post-crud-service/domain"
)

type PostCrudService interface {
	GetAllPosts(page, limit int) ([]domain.Post, error)
	GetOne(postId int) (*domain.Post, error)
	Update(postId int, title, body string) (*domain.PostUpdateResponse, error)
	Delete(postId int) error
}

func NewPostCrudUseCase(postRepo domain.PostCrudRepository) PostCrudService {
	return &postUseCase{postRepo: postRepo}
}

type postUseCase struct {
	postRepo domain.PostCrudRepository
}

func (p *postUseCase) GetAllPosts(page, limit int) ([]domain.Post, error) {
	posts, err := p.postRepo.GetList(page, limit)
	if err != nil {
		return nil, fmt.Errorf("failed to get all posts: %v", err)
	}
	return posts, nil
}

func (p *postUseCase) GetOne(postId int) (*domain.Post, error) {
	post, err := p.postRepo.GetOne(postId)
	if err != nil {
		return nil, fmt.Errorf("failed to get post %v", err)
	}
	return post, nil
}

func (p *postUseCase) Update(postId int, title, body string) (_ *domain.PostUpdateResponse, err error) {
	if title == "" || body == "" {
		return &domain.PostUpdateResponse{
			Success: false,
			Message: fmt.Sprintf("title and body required for update"),
		}, err
	}
	updateRes, err := p.postRepo.Update(postId, title, body)
	if err != nil {
		return &domain.PostUpdateResponse{
			Success: false,
			Message: fmt.Sprintf("failed to update post %v", err),
		}, err
	}

	return updateRes, nil
}

func (p *postUseCase) Delete(postId int) error {
	err := p.postRepo.Delete(postId)
	if err != nil {
		return fmt.Errorf("can't delete post %v", err)
	}
	return nil
}
