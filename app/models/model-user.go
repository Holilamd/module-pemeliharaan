package models

import "time"

type Userinput struct {
	Id        string `json:"id" `
	Nama      string `json:"nama" binding:"required"`
	Username  string `json:"username" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required"`
	Status    bool   `json:"status"`
	Locked    bool   `json:"locaked"`
	Csalahpwd int    `json:"csalahpwd"`
	CreatedBy string `json:"created_by"`
	UpdatedBy string `json:"updated_by"`
}

type User struct {
	Id        string
	Nama      string
	Username  string
	Email     string
	Status    bool
	Password  string
	Locked    bool
	Csalahpwd int
	CreatedBy string
	UpdatedBy string
	CreatedAt time.Time
	UpdatedAt time.Time
}
