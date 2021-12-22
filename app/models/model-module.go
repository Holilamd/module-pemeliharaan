package models

import "time"

type Moduleinput struct {
	Modulename    string `json:"module_name" binding:"required"`
	Url           string `json:"url" binding:"required"`
	Icon          string `json:"icon" binding:"required"`
	Modulecaption string `json:"module_caption" binding:"required"`
	Modulejenis   string `json:"module_jenis" binding:"required"`
	Moduleparent  string `json:"module_parentid" `
	Moduleorder   string `json:"module_order"`
	Status        bool   `json:"status"`
	CreatedBy     string `json:"created_by"`
	UpdatedBy     string `json:"updated_by"`
}

type Module struct {
	Modulename    string
	Url           string
	Icon          string
	Modulecaption string
	Modulejenis   string
	Moduleparent  string
	Moduleorder   string
	Status        bool
	CreatedBy     string
	UpdatedBy     string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
