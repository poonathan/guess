package main

import (
	"fmt"
	"math/rand"
	"time"
)

func guess(num int) bool {
	fmt.Print("What is your guess? ")
	var input int
	fmt.Scanf("%d\n", &input)

	switch {
	case input < num:
		fmt.Printf("Your guess %d is lower than what I thought.\n", input)
		return false

	case input > num:
		fmt.Printf("Your guess %d is higher than what I thought.\n", input)
		return false

	default:
		fmt.Printf("You got it right! I thought of %d.\n", num)
		return true
	}
}

func main() {
	rand.Seed(time.Now().Unix())
	num := rand.Intn(100) + 1

	hasWon := false
	for i := 0; i < 7 || !hasWon; i++ {
		hasWon = guess(num)
	}

	if !hasWon {
		fmt.Println("Better luck next time.")
	}
}
