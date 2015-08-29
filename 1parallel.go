package main

import (
	"log"
)

const limit = 1e3

func main() {
	finished := make(chan int)
	go Crawl(1, finished)

	for id := range make([]int, 1+(2*limit)) {
		<-finished
		log.Println(id, "finished")
	}

	close(finished)
}

func Crawl(id int, finished chan<- int) {
	log.Println(id, "start")
	finished <- id
	log.Println("\t", id, "answered")
	if id <= limit {
		go Crawl(2*id, finished)
		Crawl((2*id)+1, finished)
	}

	log.Println("\t", id, "last line")
}
