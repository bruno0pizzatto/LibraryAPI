package main

import "github.com/bruno0pizzatto/LibraryAPI/server"

func main() {
	server := server.NewServer()

	server.Run()
}
