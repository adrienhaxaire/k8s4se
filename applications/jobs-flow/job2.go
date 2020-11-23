package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	content, err := ioutil.ReadFile("/app/data/jobs.txt")
	if err != nil {
		log.Fatal(err)
	}
	upd := append(content, []byte(" and from job2")...)
	err = ioutil.WriteFile("/app/data/jobs.txt", upd, 0644)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Final message: %s", upd)
}
