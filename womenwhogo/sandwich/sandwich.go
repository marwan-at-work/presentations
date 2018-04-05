package main

import (
	"fmt"
	"time"
)

func main() {
	startTime := time.Now()
	pb := getPeanutButter()
	j := getJelly()
	makeSandwich(pb, j)
	fmt.Printf("Sandwich took %v to make\n", time.Since(startTime))
}

func makeSandwich(ingredientOne, ingredientTwo string) {
	fmt.Println(
		"putting together " + ingredientOne + " and " + ingredientTwo,
	)
}

func getPeanutButter() string {
	fmt.Println("Butler 1 says: I am getting the peanut butter")
	simulateWork()
	return "peanut butter"
}

func getJelly() string {
	fmt.Println("Butler 1 says: I am getting the jelly")
	simulateWork()
	return "jelly"
}

func simulateWork() {
	time.Sleep(time.Second)
}
