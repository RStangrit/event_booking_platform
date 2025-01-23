package main

import (
	"main/cmd/server"
	"main/pkg/database"
)

func main() {
	database.InitDB()
	server.LaunchServer()
}
