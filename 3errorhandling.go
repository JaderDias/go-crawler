package main

import (
	"log"
    "fmt"
)

const searchSpace = 1e3

func main() {
	channel := make(chan Response)
	to_work := []int{1}
	working := make(map[int]bool)
	for len(working) > 0 || len(to_work) > 0 {
		for len(to_work) > 0 && len(working) < 2 {
			id := to_work[0]
			to_work = to_work[1:]
			working[id] = true
			go Crawl(id, channel)
		}

		log.Println("working", working)
        select {
            case response := <-channel:
                delete(working, response.id)
                for _, link := range response.links {
                    to_work = append(to_work, link)
                }
            case <-time.After(1 * time.Second):
                log.Println("")
        }
	}

	close(channel)
}

func Crawl(id int, channel chan<- Response) {
	links := []int{}
    channel <- Response{id, fmt.Errorf("failed"), links}
	if id < searchSpace {
		links = []int{2 * id, (2 * id) + 1}
	}

	channel <- Response{id, nil, links}
}

type Response struct {
	id    int
    err   Error
	links []int
}
