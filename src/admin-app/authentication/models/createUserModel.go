package models

type BFFCreateUserRequest struct {
	Name        int `json:"name"`
	Username    int `json:"username"`
	PhoneNumber int `json:"phoneNumber"`
	Password    int `json:"password"`
	PanCard     int `json:"panCard"`
	Email       int `json:"email"`
}

type BFFCreateUserResponse struct {
	Name        string `json:"name"`
	Username    string `json:"username"`
	PhoneNumber string `json:"phoneNumber"`
	Password    string `json:"password"`
	PanCard     string `json:"panCard"`
	Email       string `json:"email"`
	Message     string `json:"message"`
}
