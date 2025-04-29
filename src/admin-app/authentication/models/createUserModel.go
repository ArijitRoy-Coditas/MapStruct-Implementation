package models

type BFFCreateUserRequest struct {
	Name        int         `json:"name"`        // int → will convert to string
	Username    string      `json:"username"`    // string → int
	PhoneNumber interface{} `json:"phoneNumber"` // float64 → uint64
	Password    string      `json:"password"`    // same
	PanCard     float64     `json:"panCard"`     // float64 → int64
	Email       int         `json:"email"`       // int → string
}

type BFFCreateUserResponse struct {
	Name        string `json:"name"`        // from int → string
	Username    int    `json:"username"`    // from string → int
	PhoneNumber uint64 `json:"phoneNumber"` // from interface(float64) → uint64
	Password    string `json:"password"`    // same
	PanCard     int64  `json:"panCard"`     // from float64 → int64
	Email       string `json:"email"`       // from int → string
}
