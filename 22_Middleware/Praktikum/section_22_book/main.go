package main

import (
	"section_22_book/config"
	"section_22_book/route"
)

func main() {
	config.InitDB()
	e := route.New()
	e.Start(":8000")
}
