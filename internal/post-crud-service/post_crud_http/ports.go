package post_crud_http

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"test-project-iman/internal/post-crud-service/app"
	"test-project-iman/internal/post-crud-service/domain"
)

type PostCrudController struct {
	postUseCase app.PostCrudService
}

func NewPostCrudController(postUseCase app.PostCrudService) *PostCrudController {
	return &PostCrudController{postUseCase: postUseCase}
}

type PostIDRequest struct {
	PostID int `json:"postId"`
}

func (p *PostCrudController) GetAll(w http.ResponseWriter, r *http.Request) {
	var postRes domain.Post
	err := json.NewDecoder(r.Body).Decode(&postRes)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
	}
	posts, err := p.postUseCase.GetAllPosts()
	if err != nil {
		http.Error(w, "posts not found", http.StatusNotFound)
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(posts)
}

func (p *PostCrudController) GetOne(w http.ResponseWriter, r *http.Request) {
	var requestPayload PostIDRequest
	err := json.NewDecoder(r.Body).Decode(&requestPayload)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	postId := requestPayload.PostID

	post, err := p.postUseCase.GetOne(postId)
	if err != nil {
		http.Error(w, "Post not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(post)
}

func (p *PostCrudController) Update(w http.ResponseWriter, r *http.Request) {
	var updateRequest struct {
		PostID int    `json:"postId"`
		Title  string `json:"title"`
		Body   string `json:"body"`
	}

	err := json.NewDecoder(r.Body).Decode(&updateRequest)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	postId := updateRequest.PostID
	title := updateRequest.Title
	body := updateRequest.Body

	if title == "" || body == "" {
		http.Error(w, "Title and body are required for update", http.StatusBadRequest)
		return
	}
	updateRes, err := p.postUseCase.Update(postId, title, body)
	if err != nil {
		log.Println("Failed: ", err.Error())
		http.Error(w, fmt.Sprintf("Failed to update post: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updateRes)
}

func (p *PostCrudController) Delete(w http.ResponseWriter, r *http.Request) {

	var deleteRequest struct {
		PostID int `json:"postId"`
	}
	err := json.NewDecoder(r.Body).Decode(&deleteRequest)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
	}

	postId := deleteRequest.PostID

	err = p.postUseCase.Delete(postId)
	if err != nil {
		http.Error(w, "Couldn't delete", http.StatusNotFound)
	}

	var str = "Successfully deleted"

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(str)
}
