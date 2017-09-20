package main

import (
	"fmt"
	"time"
)

func init() {
}

func main() {
	fmt.Println("porgram started")
	s := time.Now()
	makeSandwich()
	fmt.Println(time.Since(s))
}

// START OMIT
func makeSandwich() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	go getPeanutButter(ch1)
	go getJelly(ch2)

	pb := <-ch1
	j := <-ch2
	fmt.Println("putting together " + pb + " and " + j)
}

func getPeanutButter(ch chan string) {
	time.Sleep(time.Second)
	ch <- "peanut butter"
}

func getJelly(ch chan string) {
	time.Sleep(time.Second)
	ch <- "jelly"
}

// END OMIT
