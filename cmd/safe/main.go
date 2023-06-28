package main

import (
	"fmt"
	"log"

	"github.com/rur0/safe"
)

func main() {
	bbs, err := safe.Parse("/home/ruro/bals_e")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(bbs.Sum())
}
