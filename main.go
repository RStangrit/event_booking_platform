package main

import (
	"main/cmd/server"   //launch server
	"main/pkg/database" //launch db
	"main/pkg/util"     //get env variables
)

func main() {
	dbType := util.GetEnvVariable("DATABASE_TYPE")
	database.InitDB(dbType)
	server.LaunchServer()
}
