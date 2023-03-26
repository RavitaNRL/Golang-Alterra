package models

type User struct {
	ID       int    `json:"ID" form:"id"`
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type UserResponse struct {
	ID    int    `json:"id" form:"id"`
	Name  string `json:"name" form:"name"`
	Email string `json:"email" form:"email"`
	Token string `json:"token" form:"token"`
}

// type Buku struct {
// 	ID       int    `json:"ID" form:"id"`
// 	Judul    string `json:"judul" form:"judul"`
// 	Penulis  string `json:"penulis" form:"penulis"`
// 	Penerbit string `json:"penerbit" form:"penerbit"`
// }
