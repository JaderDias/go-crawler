package main

import (
	"log"
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
		response := <-channel
		delete(working, response.id)
		for _, link := range response.links {
			to_work = append(to_work, link)
		}
	}

	close(channel)
}

func Crawl(id int, channel chan<- Response) {
	links := []int{}
	if id < searchSpace {
		links = []int{2 * id, (2 * id) + 1}
	}

	channel <- Response{id, links}
}

type Response struct {
	id    int
	links []int
}
