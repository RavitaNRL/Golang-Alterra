package main

import (
	"section23/config"
	"section23/route"
)

func main() {
	config.InitDB()
	e := route.New()
	e.Start(":8000")
}
