package main

import (
	"flag"
	"log"
)

var param = flag.String("param", "value", "param")

func main() {
	flag.Parse()
	log.Printf("hello")
}
