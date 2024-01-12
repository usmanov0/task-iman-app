package app

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"test-project-iman/internal/post-crud-service/domain"
	"testing"
)

type MockPostCrudRepository struct {
	mock.Mock
}

func (m *MockPostCrudRepository) GetList() ([]domain.Post, error) {
	args := m.Called()
	return args.Get(0).([]domain.Post), args.Error(1)
}

func (m *MockPostCrudRepository) GetOne(postId int) (*domain.Post, error) {
	args := m.Called(postId)
	return args.Get(0).(*domain.Post), args.Error(1)
}

func (m *MockPostCrudRepository) Update(postId int, title, body string) (*domain.PostUpdateResponse, error) {
	args := m.Called(postId, title, body)
	return args.Get(0).(*domain.PostUpdateResponse), args.Error(1)
}

func (m *MockPostCrudRepository) Delete(postId int) error {
	args := m.Called(postId)
	return args.Error(0)
}

func TestGetAllPosts(t *testing.T) {
	mockRepo := &MockPostCrudRepository{}
	postService := NewPostCrudUseCase(mockRepo)

	// Mock GetList method to return dummy posts
	mockRepo.On("GetList").Return([]domain.Post{{Id: 1, Title: "Post 1"}, {Id: 2, Title: "Post 2"}}, nil).Once()

	// Call the GetAllPosts method
	posts, err := postService.GetAllPosts()

	// Assertions
	assert.NoError(t, err)
	assert.Len(t, posts, 2)
	assert.Equal(t, "Post 1", posts[0].Title)
	assert.Equal(t, "Post 2", posts[1].Title)
	mockRepo.AssertExpectations(t)
}

func TestGetOne(t *testing.T) {
	mockRepo := &MockPostCrudRepository{}
	postService := NewPostCrudUseCase(mockRepo)

	mockRepo.On("GetOne", 1).Return(&domain.Post{Id: 1, Title: "Post 1"}, nil).Once()

	post, err := postService.GetOne(1)

	// Assertions
	assert.NoError(t, err)
	assert.NotNil(t, post)
	assert.Equal(t, "Post 1", post.Title)
	mockRepo.AssertExpectations(t)
}

func TestUpdate(t *testing.T) {
	mockRepo := &MockPostCrudRepository{}
	postService := NewPostCrudUseCase(mockRepo)

	t.Run("Successfully update post", func(t *testing.T) {
		// Mock Update method to return a success response
		mockRepo.On("Update", 1, "Updated Title", "Updated Body").Return(
			&domain.PostUpdateResponse{Success: true, Message: "Post updated successfully"}, nil,
		).Once()

		response, err := postService.Update(1, "Updated Title", "Updated Body")

		assert.NoError(t, err)
		assert.NotNil(t, response)
		assert.True(t, response.Success)
		assert.Equal(t, "Post updated successfully", response.Message)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Update with empty title and body", func(t *testing.T) {
		response, err := postService.Update(1, "", "")

		assert.Error(t, err)
		assert.NotNil(t, response)
		assert.False(t, response.Success)
		assert.Equal(t, "title and body required for update", response.Message)
	})

	t.Run("Error updating post", func(t *testing.T) {
		mockRepo.On("Update", 1, "Updated Title", "Updated Body").Return(
			nil, errors.New("mock error"),
		).Once()

		response, err := postService.Update(1, "Updated Title", "Updated Body")

		assert.Error(t, err)
		assert.NotNil(t, response)
		assert.False(t, response.Success)
		assert.Equal(t, "failed to update post mock error", response.Message)
		mockRepo.AssertExpectations(t)
	})
}

func TestDelete(t *testing.T) {
	mockRepo := &MockPostCrudRepository{}
	postService := NewPostCrudUseCase(mockRepo)

	mockRepo.On("Delete", 1).Return(nil).Once()

	err := postService.Delete(1)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}
