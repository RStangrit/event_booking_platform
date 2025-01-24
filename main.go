package main

import (
	"main/cmd/server"   //launch server
	"main/pkg/database" //launch db
	"main/pkg/util"     //get env variables
)

func main() {
	dbOperator := util.GetEnvVariable("DATABASE_OPERATOR")
	database.InitDB(dbOperator)
	server.LaunchServer()
}
