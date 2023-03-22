package main

import (
	"prio2/config"
	"prio2/route"
)

func main() {
	config.InitDB()
	e := route.New()
	e.Start(":8000")
}
