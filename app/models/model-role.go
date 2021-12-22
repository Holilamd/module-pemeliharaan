package models

import "time"

type Roleinput struct {
	Rolename  string `json:"role_name" binding:"required"`
	Status    bool   `json:"status"`
	CreatedBy string `json:"created_by"`
	UpdatedBy string `json:"updated_by"`
}

type Role struct {
	Rolename  string
	Status    bool
	CreatedBy string
	UpdatedBy string
	CreatedAt time.Time
	UpdatedAt time.Time
}
