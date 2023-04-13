package main

import (
	"flag"

	"github.com/farrelahmady/my-gram-hacktiv8/database"
)

func main() {
	database.InitDB()
	HandleArgs()

}

func HandleArgs() {
	flag.Parse()
	args := flag.Args()
	_ = args
}
