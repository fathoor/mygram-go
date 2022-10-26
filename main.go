package main

import (
	"github.com/fathoor/mygram-go/database"
	"github.com/fathoor/mygram-go/router"
)

func main() {
	database.StartDB()
	r := router.StartApp()
	r.Run(database.APP_HOST + ":" + database.APP_PORT)
}
