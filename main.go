package main

import (
	"fmt"
)

func main() {
	fmt.Println("\t\tWelcome to Video Streaming Platform\t\t")
	DisplayMenu()
	active := true
	active, userInput := ValidateInput()
	for active {
		switch userInput {
		case 1:
			SearchShow()
		case 2:
			FetchTopShows()
		case 3:
			FetchMostRecentShows()
		case 4:
			FetchMostFrequentlyWatchedShows()
		}
		DisplayMenu()
		active, userInput = ValidateInput()
	}
	Exit()
}
