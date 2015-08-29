package main

import (
	"log"
    "time"
)

const searchSpace = 1e3
const max_parallelism = 2

type Status struct {
    id int
    finished bool
}

func main() {
	channel := make(chan Status, max_parallelism)
    working := []int{}
	go Crawl(1, channel)
	for _ = range make([]int, searchSpace * 2) {
        status := <-channel
        if status.finished {
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
}

func Crawl(id int, channel chan<- Status) {
	channel <- Status{id, false}
    time.Sleep(1)
	channel <- Status{id, true}
    nextId := 2*id
	if nextId <= searchSpace {
		go Crawl(nextId, channel)
    }

    nextId++
    if nextId <= searchSpace {
		Crawl(nextId, channel)
	}
}
