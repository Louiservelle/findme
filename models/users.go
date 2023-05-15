package models

type User struct {
	ID           int    `json:"user_id"`
	Name         string `json:"name"`
	Password     string `json:"password"`
	Email        string `json:"email"`
	Organization string `json:"organisation"`
	Phone        string `json:"Phone"`
	MissionID    int    `json:"mission_id"`
}
