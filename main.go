package main

import "github.com/mattdotmatt/bigstar/server"

func main() {
	server.Start("localhost", 8181, "./public/data/db.json")
}
