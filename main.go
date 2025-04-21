package main

import (
	"fmt"
)

func getName() string {
	name := ""
	fmt.Println("Welcome to my Casino...")
	fmt.Printf("Enter your name: ")
	_, err := fmt.Scanln(&name)
	if err != nil {
		return ""
	}
	fmt.Printf("Welcome %s, let's play!\n", name)
	return name
}

func getBet(balance uint) uint {
	var bet uint
	for true {
		fmt.Printf("Enter your bet, or 0 to quit (Balance = $%d): ", balance)
		fmt.Scan(&bet)
		if bet > balance {
			fmt.Println("Bet cannot be larger than balance.")
		} else {
			break
		}
	}
}

func main() {
	balance := 200
	getName()
}
