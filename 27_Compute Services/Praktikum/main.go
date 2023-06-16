package main

import (
	"section22/config"
	"section22/route"
)

func main() {
	config.InitDB()
	e := route.New()
	e.Start(":8000")
}
