package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	message := []byte("Writing from job1")
	err := ioutil.WriteFile("/app/data/jobs.txt", message, 0644)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Final message: %s", message)
}
