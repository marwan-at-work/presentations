package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("program started")
	s := time.Now()
	makeSandwich()
	fmt.Println(time.Since(s))
}

// START OMIT
func makeSandwich() {
	pb := getPeanutButter()
	j := getJelly()
	fmt.Println(
		"putting together " + pb + " and " + j,
	)
}

func getPeanutButter() string {
	time.Sleep(time.Second)
	return "peanut butter"
}

func getJelly() string {
	time.Sleep(time.Second)
	return "jelly"
}

// END OMIT
