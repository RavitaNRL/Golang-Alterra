package models

type Buku struct {
	ID       int    `json:"ID" form:"id"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Judul    string `json:"judul" form:"judul"`
	Penulis  string `json:"penulis" form:"penulis"`
	Penerbit string `json:"penerbit" form:"penerbit"`
	Token    string `json:"token" form:"token"`
}
