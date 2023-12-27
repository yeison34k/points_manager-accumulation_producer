package domain

type Point struct {
	ID    string  `json:"id"`
	User  string  `json:"user"`
	Name  string  `json:"name"`
	Total float32 `json:"total"`
}

type Response struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}
