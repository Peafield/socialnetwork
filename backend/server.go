package main

import db "socialnetwork/pkg/db/sqlite"

func main() {
	db.InitialiseDatabase("./pkg/db", "socialNetworkDB")
}
