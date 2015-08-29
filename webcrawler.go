package main

import (
	"log"
)

func main() {
    semaphore := make(chan int)
	finished := make(chan int)
    Crawl(1, semaphore, finished)

    semaphore<-1
    semaphore<-1
	for id := range finished  {
        log.Println(id, "finished")
        semaphore<-1
	}

    close(finished)
    close(semaphore)
}

func Crawl(id int, semaphore, finished chan int) {
	log.Println(id, "ready")
    // <-semaphore
	log.Println(id, "start")
	if id < 3 {
		go Crawl(2*id, semaphore, finished)
		go Crawl((2*id)+1, semaphore, finished)
	}

    finished <- id
	log.Println(id, "last line")
}
