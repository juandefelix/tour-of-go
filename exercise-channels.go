package main

import "fmt"

func sum(s []int, channel chan int) {
	sum := 0
	for _, value := range(s) {
		sum += value
	}

	channel <- sum
}

func main() {
	integers := []int{1,2,3,4,5,6}
	channel := make(chan int)

	go sum(integers[len(integers)/2:], channel)
	go sum(integers[:len(integers)/2], channel)

	sum1, sum2 := <-channel, <-channel

	fmt.Println(sum1, sum2)
}
