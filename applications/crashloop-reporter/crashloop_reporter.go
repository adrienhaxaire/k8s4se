package main

import (
	"flag"
	"log"
)

func main() {
	var arg *string
	arg = flag.String("arg1", "toto", "the argument")
	flag.Parse()
	println(*arg)
	log.Printf("argument: %s", *arg)

}
