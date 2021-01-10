package main

import (
	"github.com/fregdee/gols"
	"log"
)

func main() {
	err := gols.Run()
	if err != nil {
		log.Fatalln(err)
	}
}
