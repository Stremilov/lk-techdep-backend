package project

type Project struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Status      string `json:"status"`
	Role        string `json:"role"`
	Description string `json:"description"`
}
