package main

import (
	"io/ioutil"
	"log"
)

func main() {
	content, err := ioutil.ReadFile("data/jobs.txt")
	if err != nil {
		log.Fatal(err)
	}
	upd := append(content, []byte(" and from job2")...)
	err = ioutil.WriteFile("data/jobs.txt", upd, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
