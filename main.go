package main

import (
	"main/cmd/server"
	"main/pkg/database"
)

func main() {
	server.LaunchServer()

	database.InitDB()
}
