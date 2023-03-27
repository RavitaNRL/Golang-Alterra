package main

import (
	"section22_buku/config"
	"section22_buku/route"
)

func main() {
	config.InitDB()
	e := route.New()
	e.Start(":8000")
}
