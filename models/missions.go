package models

type Mission struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Skills      string `json:"skills"`
	UserId      int    `json:"user_id"`
}
