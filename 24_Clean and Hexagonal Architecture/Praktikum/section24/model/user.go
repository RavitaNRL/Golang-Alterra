package model

import "gorm.io/gorm"

// membuat sturct yang akan terhubung dengan tabel di database dan terisi otomatis oleh gorm.
type User struct {
	*gorm.Model

	ID       int    `json:"id" form:"id"`
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Token    string `json:"token" form:"token"`
}
