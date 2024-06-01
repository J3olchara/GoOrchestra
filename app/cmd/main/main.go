package main

import (
	"github.com/J3olchara/GoOrchestra/app/server/api"
	"github.com/J3olchara/GoOrchestra/app/server/db"
	"log"
	"os"
)

func main() {
	log.SetFlags(log.Ltime | log.Ldate)
	api.SVR = api.NewServer(os.Getenv("APP_HOST"), os.Getenv("APP_PORT"))
	db.DB = db.NewDatabase()
	err := db.DB.PrepareDatabase()
	if err != nil {
		log.Fatal(err)
	}
	api.SVR.StartServer()
	return
}
