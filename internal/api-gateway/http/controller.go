package http

import (
	"encoding/json"
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
	err := c.service.FetchPosts()

	if err != nil {
		http.Error(w, "error fetching data "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Data fetched successfully and saved to database")
}

func (c *HttpServer) GetPosts(w http.ResponseWriter, r *http.Request) {
	posts, err := c.service.GetList()
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

	updateResponse, err := c.service.Update(id, title, body)
	if err != nil {
		http.Error(w, "Failed to update post", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updateResponse)
}

func (c *HttpServer) Delete(w http.ResponseWriter, _ *http.Request) {
	var id int

	err := c.service.Delete(id)
	if err != nil {
		http.Error(w, "Failed to delete post", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Deleted successfully")
}
