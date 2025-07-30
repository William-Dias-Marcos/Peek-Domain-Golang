package main

import (
	"log"
	"os"

	"github.com/William-Dias-Marcos/Peek-Domain-Golang/app"
)

func main() {
	application := app.Generate()
	if err := application.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}