package main

import (
	"flag"
	"log"
	"strings"

	"github.com/farrelahmady/my-gram-hacktiv8/database"
	"github.com/farrelahmady/my-gram-hacktiv8/migrations"
)

func main() {
	database.InitDB()
	HandleArgs()

}

func HandleArgs() {
	flag.Parse()
	args := flag.Args()
	for _, arg := range args {
		command := strings.Split(arg, "=")[0]
		var params []string
		if len(strings.Split(arg, "=")) > 1 {
			params = strings.Split(strings.Split(arg, "=")[1], ",")
		}
		// log.Println(command, params)
		switch command {
		case "migrate":
			db := database.GetDB()
			migrations.Execute(db, params...)
		case "seed":
			log.Println("seed")
		}

	}
}
