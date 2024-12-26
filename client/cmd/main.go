package main

import (
	"log"

	"github.com/jchapman63/neo-gkmn/client"
)

func main() {
	g, err := client.NewGame()
	if err != nil {
		log.Fatal(err)
	}
	if err := client.Run(g); err != nil {
		log.Fatal(err)
	}
}
