package app

import (
	"errors"
	"test-project-iman/internal/post-collector-service/domain"
	"testing"
)

type MockPostRepository struct {
	SaveFunc func(post *domain.Post) (int, error)
}

type MockPostProviderRepository struct {
	FetchPostsFunc func(page string) ([]domain.Post, error)
}

func (m *MockPostRepository) Save(post *domain.Post) (int, error) {
	if m.SaveFunc != nil {
		return m.SaveFunc(post)
	}
	return 0, errors.New("SaveFunc not implemented")
}

func (m *MockPostProviderRepository) FetchPosts(page string) ([]domain.Post, error) {
	if m.FetchPostsFunc != nil {
		return m.FetchPostsFunc(page)
	}
	return nil, errors.New("FetchPostsFunc not implemented")
}

// TODO optimize
func TestPostCollectorService_CollectPosts(t *testing.T) {
	mockPostRepo := &MockPostRepository{
		SaveFunc: func(post *domain.Post) (int, error) {
			return 1, nil
		},
	}
	mockPostProviderRepo := &MockPostProviderRepository{
		FetchPostsFunc: func(page string) ([]domain.Post, error) {
			return []domain.Post{{Id: 1, UserId: 1, Title: "Test", Body: "Test Body"}}, nil
		},
	}

	postService := NewPostService(mockPostRepo, mockPostProviderRepo)

	err := postService.CollectPosts()

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}

// TODO optimize
func TestPostCollectorService_CollectPosts_ErrorFetchingPosts(t *testing.T) {
	mockPostRepo := &MockPostRepository{
		SaveFunc: func(post *domain.Post) (int, error) {
			return 0, errors.New("save error")
		},
	}
	mockPostProviderRepo := &MockPostProviderRepository{
		FetchPostsFunc: func(page string) ([]domain.Post, error) {
			return nil, errors.New("fetch error")
		},
	}

	postService := NewPostService(mockPostRepo, mockPostProviderRepo)

	err := postService.CollectPosts()

	if err == nil {
		t.Error("Expected an error but got nil")
	}
}
