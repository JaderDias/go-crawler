package main

import (
	"log"
)

func main() {
	channel := make(chan Response)
	to_work := []int{1}
	working := make(map[int]bool)
	for len(working) > 0 || len(to_work) > 0 {
		for len(to_work) > 0 && len(working) < 2 {
			id := to_work[0]
			to_work = to_work[1:]
			working[id] = true
			go func() {
				links := []int{}
				if id < 1e3 {
					links = []int{2 * id, (2 * id) + 1}
				}

				channel <- Response{id, links}
			}()
		}

		log.Println("working", working)
		response := <-channel
		delete(working, response.id)
		for _, link := range response.links {
			to_work = append(to_work, link)
		}
	}
}

type Response struct {
	id    int
	links []int
}
