package main

import "fmt"

func main() {
	c := generate(4, 10)

	d1 := divide(c)
	d2 := divide(c)

	x := funnel(d1, d2)
	fmt.Println(<-x)
	fmt.Println(<-x)
	fmt.Println("Finished!")
}

func generate(numbers ...int) chan int {
	channel := make(chan int)

	go func() {
		for _, n := range numbers {
			channel <- n
		}
		close(channel)
	}()
	return channel
}

func divide(input chan int) chan int {
	channel := make(chan int)

	go func() {
		for number := range input {
			channel <- number / 2
		}
		close(channel)
	}()

	return channel
}

func funnel(channel1, channel2 <-chan int) <-chan int {
	channel := make(chan int)

	go func() {
		for {
			channel <- <-channel1
		}
	}()

	go func() {
		for {
			channel <- <-channel2
		}
	}()

	return channel
}
