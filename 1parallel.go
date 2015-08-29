package main

import (
	"log"
)

const searchSpace = 1e3

func main() {
	channel := make(chan Response)
	go Crawl(1, channel)
	working := IntSlice([]int{1})
	for len(working) > 0 {
		log.Println("working", working)
		response := <-channel
		working = working.FilterOut(response.id)
		for _, link := range response.links {
			go Crawl(link, channel)
			working = append(working, link)
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

type IntSlice []int

func (slice IntSlice) FilterOut(id int) []int {
	temp := []int{}
	for _, v := range slice {
		if v != id {
			temp = append(temp, v)
		}
	}

	return temp
}
