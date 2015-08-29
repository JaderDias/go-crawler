package main

import (
	"log"
    "time"
)

const searchSpace = 1e3

type Status struct {
    id int
    finished bool
}

func main() {
	channel := make(chan Status)
    semaphore := make(chan int, 2)
    semaphore<-1
    semaphore<-1
    working := []int{}
	go Crawl(1, channel, semaphore)
	for _ = range make([]int, searchSpace * 2) {
        status := <-channel
        if status.finished {
            semaphore<- 1
            temp := []int{}
            for _, v := range working {
                if v != status.id {
                    temp = append(temp, v)
                }
            }

            working = temp
        } else {
            working = append(working, status.id)
        }

        log.Println("working", working)
	}

	close(channel)
    close(semaphore)
}

func Crawl(id int, channel chan<- Status, semaphore <-chan int) {
    <-semaphore
	channel<- Status{id, false}
    time.Sleep(1)
	channel<- Status{id, true}
    nextId := 2*id
	if nextId <= searchSpace {
		go Crawl(nextId, channel, semaphore)
    }

    nextId++
    if nextId <= searchSpace {
		Crawl(nextId, channel, semaphore)
	}
}
