package main

import (
	"kafkatool/internal/server"
	"log"
)

func main() {
	app := server.GetApp()
	if err := app.Run(); err != nil {
		log.Fatalln(err)
	}
}
