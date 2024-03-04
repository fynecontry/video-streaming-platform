package main

import (
	"fmt"
	"strings"
)

func SearchShow() {
	titles := []string{"king", "barbie", "inception", "speed", "tenet", "maverick", "code8", "frozen", "jumanji"}
	fmt.Println("Enter show's title: ")
	var query string
	fmt.Scan(&query)
	query = strings.ToLower(query)
	for _, title := range titles {
		if IsAnagram(query, title) {
			fmt.Printf("%s found.\n", title)
			return
		}
	}
	fmt.Printf("%s not found.\n", query)
}

func FetchTopShows() {
	ratings1 := [][]interface{}{{"King", 11}, {"Barbie", 23}, {"Inception", 32}, {"Speed", 38}}
	ratings2 := [][]interface{}{{"Tenet", 1}, {"Maverick", 44}}
	ratings3 := [][]interface{}{{"Code8", 23}, {"Frozen", 57}, {"Jumanji", 99}}
	nasa := createLinkedList(ratings1)
	emea := createLinkedList(ratings2)
	apac := createLinkedList(ratings3)
	showList := []*LinkedListNode{nasa, emea, apac}
	display(MergeKSortedLists(showList))
}

func FetchMostRecentShows() {
	cache := &LRUCache{capacity: 3, cache: make(map[string]*LinkedListNode)}
	fmt.Println("The most recently watched titles are: (Show, rank)")
	cache.Set("Barbie", 23)
	cache.Print()

	cache.Set("King", 11)
	cache.Print()

	cache.Set("Inception", 32)
	cache.Print()

	cache.Set("Code8", 23)
	cache.Print()

	cache.Set("Tenet", 1)
	cache.Print()

	cache.Get("Barbie")
	cache.Print()
}

func Exit() {
	fmt.Println("Goodbye!")
}
