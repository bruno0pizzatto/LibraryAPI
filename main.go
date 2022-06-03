package main

import (
	"github.com/bruno0pizzatto/LibraryAPI/database"
	"github.com/bruno0pizzatto/LibraryAPI/database/migrations"
	"github.com/bruno0pizzatto/LibraryAPI/server"
)

func main() {
	database.StartDB()
	server := server.NewServer()
	migrations.RunMMigrations(database.GetDataBase())
	server.Run()
}
