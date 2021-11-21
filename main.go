package main

import (
	"fmt"
	"ip-searcher/app"
	"log"
	"os"
)

func main() {
	fmt.Println("Starting App ...")

	application := app.Generate()
	if error := application.Run(os.Args); error != nil {
		log.Fatal(error)
	}

}
