package domain

type Post struct {
	Id     int    `json:"id"`
	UserId int    `json:"userId"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

type PostUpdateResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}
