package main

import (
	"section_22/config"
	"section_22/route"
)

func main() {
	config.InitDB()
	e := route.New()
	e.Start(":8000")
}
