package model

type Post struct {
	ID     int    `json:"id"`
	UserId int    `json:"user_id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

type Posts struct {
	Post
}

type UpdateRequest struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

type PostUpdateResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}
