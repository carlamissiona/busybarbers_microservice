package main

import (
	"carlamissiona/golang-barbers/bootstrap"
	"log"
	"os"
)

func main() {

	if len(os.Args) == 2 {
		if os.Args[1] == "monolith" {

			app := bootstrap.NewApplication("monolith")
			log.Fatal(app.Listen(":5000"))

		}
		if os.Args[1] == "api" {

			app := bootstrap.NewApplication("api")
			log.Fatal(app.Listen(":8000"))

		}
	}

}
