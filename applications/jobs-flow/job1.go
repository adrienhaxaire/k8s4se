package main

import (
	"io/ioutil"
	"log"
)

func main() {
	message := []byte("Writing from job1")
	err := ioutil.WriteFile("data/jobs.txt", message, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
