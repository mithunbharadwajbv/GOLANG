package main

import "fmt"

func main() {

	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			ch1 <- i
		}
		close(ch1)
	}()

	go worker(ch1, ch2)
	go worker(ch1, ch2)

	for val := range ch2 {
		fmt.Println(val)
	}

}

func worker(ch1 chan int, ch2 chan int) {
	for val := range ch1 {
		ch2 <- FibonacciRecursion(val)
	}
}

func FibonacciRecursion(n int) int {
	if n <= 1 {
		return n
	}
	return FibonacciRecursion(n-1) + FibonacciRecursion(n-2)
}
