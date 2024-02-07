package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"test-project-iman/internal/api-gateway/app"
	"test-project-iman/internal/api-gateway/model"
)

type HttpServer struct {
	service app.Service
}

func NewController(service app.Service) *HttpServer {
	return &HttpServer{
		service: service,
	}
}

func (c *HttpServer) CollectPostsHandler(w http.ResponseWriter, r *http.Request) {
	err := c.service.CollectPosts()

	if err != nil {
		http.Error(w, "error fetching data "+err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"message": fmt.Sprint("Posts collected successfully and saved to database"),
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func (c *HttpServer) GetPosts(w http.ResponseWriter, r *http.Request) {
	var page, limit int
	posts, err := c.service.GetList(page, limit)
	if err != nil {
		http.Error(w, "Failed to get posts"+err.Error(), http.StatusInternalServerError)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(posts)
}

func (c *HttpServer) GetPost(w http.ResponseWriter, r *http.Request) {
	idStr := r.FormValue("id")
	if idStr == "" {
		http.Error(w, "Missing id parameter", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid id parameter: "+err.Error(), http.StatusBadRequest)
		return
	}

	post, err := c.service.GetById(id)
	if err != nil {
		http.Error(w, "Failed to get post"+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(post)
}

func (c *HttpServer) Update(w http.ResponseWriter, r *http.Request) {
	var updateRequest model.UpdateRequest

	err := json.NewDecoder(r.Body).Decode(&updateRequest)
	if err != nil {
		http.Error(w, "Failed to decode JSON: "+err.Error(), http.StatusBadRequest)
		return
	}
	id := updateRequest.ID
	title := updateRequest.Title
	body := updateRequest.Body

	err = c.service.Update(id, title, body)
	if err != nil {
		http.Error(w, "Failed to update post", http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"message": fmt.Sprint("Post successfully updated"),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func (c *HttpServer) Delete(w http.ResponseWriter, _ *http.Request) {
	var id int

	err := c.service.Delete(id)
	if err != nil {
		http.Error(w, "Failed to delete post", http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"message": fmt.Sprint("Post successfully deleted"),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
