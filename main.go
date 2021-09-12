package main

import (
	"log"
	"os"
	"star-wars-client/app"
)

func main() {
	app := app.CriarCli()

	err := app.Run(os.Args)

	if err != nil {
		log.Fatalln("Error: ", err)
	}
}
