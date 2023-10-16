package main

import (
	"kafkatool/internal/server"
)

func main() {
	app := server.GetApp()
	app.Run()
}
