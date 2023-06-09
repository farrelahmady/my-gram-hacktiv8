package main

import (
	"flag"

	_ "github.com/farrelahmady/my-gram-hacktiv8/docs"

	"github.com/farrelahmady/my-gram-hacktiv8/database"
	"github.com/farrelahmady/my-gram-hacktiv8/router"
)

func main() {
	database.InitDB()
	HandleArgs()

	router.StartApp().Run(":8090")
}

func HandleArgs() {
	flag.Parse()
	args := flag.Args()
	_ = args
}
