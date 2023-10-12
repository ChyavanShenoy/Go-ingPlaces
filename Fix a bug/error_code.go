package main

import "fmt"

func main() {
	messagesFromDoris := []string{
		"you doing anything later??",
		"Did you get my last message?",
		"Don't forget we are meeting at 4pm tomorrow at cafe",
		"Ping me when you are free",
	}

	numMessages := float64(len(messagesFromDoris))
	costPerMessage := 0.02

	// Don't change the code above this line.

	totalCost := numMessages + costPerMessage

	// Don't change the code below this line.

	fmt.Println("Total cost is", totalCost)
}
