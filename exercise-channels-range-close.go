package main

import "fmt"

func fibonacci(index int, channel chan int) {
	x, y := 0, 1
	for i := 0; i < index; i++ {
		channel <- x
		x, y = y, x + y
	}

	close(channel)
}

func main() {
	channel := make(chan int, 10)
	go fibonacci(cap(channel), channel)

	for value:= range channel {
		fmt.Println(value)
	}


	// You can chec when the channel has been closed this way:
	// for {
	// 	value, ok := <- channel
	// 	if !ok {
	// 		break
	// 	}
	// 	fmt.Println(value)
	// }
}
