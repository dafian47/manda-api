package main

import (
	"github.com/dafian47/manda-api/config"
	"github.com/dafian47/manda-api/router"
	"log"
	"os"
	"runtime"
)

var databaseUrl string
var port string

func init() {

	// Verbose logging with file name and line number
	log.SetFlags(log.Lshortfile)

	// Use all CPU cores
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {

	databaseUrl = os.Getenv("DATABASE_URL")
	port = os.Getenv("PORT")

	db := config.InitDB(databaseUrl)
	r := router.InitRouter(db)

	r.Run(":" + port)
}
