package main

import "github.com/mattdotmatt/bigstar/server"

func main() {
	server.Start(8181, "./web/data/db.json")
}
