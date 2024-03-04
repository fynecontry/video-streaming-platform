package main

import "fmt"

func DisplayMenu() {
	fmt.Println("1. Search for a show")
	fmt.Println("2. Display Top shows")
	fmt.Println("3. Display recently watched")
	fmt.Println("4. Exit")
	fmt.Print("Select an option: ")
}

func ValidateInput() (bool, int) {
	active := true
	var choice int
	fmt.Scan(&choice)
	for choice < 1 || choice > 4 {
		fmt.Printf("Invalid input (%d)!\n", choice)
		DisplayMenu()
		fmt.Scan(&choice)
	}
	if choice == 4 {
		active = false
	}
	return active, choice
}

func IsAnagram(s, t string) bool {
	if len(s) != len(t) {
		return false
	}

	count := make([]int, 26)

	for i := 0; i < len(s); i++ {
		count[int(s[i]-'a')]++
		count[int(t[i]-'a')]--
	}

	for _, v := range count {
		if v != 0 {
			return false
		}
	}

	return true
}
