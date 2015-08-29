package main

import (
	"log"
	"time"
)

const limit = 1e3

func main() {
	semaphore := make(chan int)
	finished := make(chan int)
	go Crawl(1, semaphore, finished)

	for id := range make([]int, 1+(2*limit)) {
		<-finished
		log.Println(id, "finished")
	}

	close(semaphore)
	close(finished)
}

func Crawl(id int, semaphore <-chan int, finished chan<- int) {
	log.Println(id, "start")
	finished <- id
	log.Println("\t", id, "answered")
	if id <= limit {
		go Crawl(2*id, semaphore, finished)
		time.Sleep(1)
		Crawl((2*id)+1, semaphore, finished)
	}

	log.Println("\t", id, "last line")
}
