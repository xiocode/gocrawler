package main

import "fmt"

func shunt(input <-chan int, output chan<- int) {
	var (
		i	int
		in	= input
		out	chan<- int
	)
	for {
		select {
		case i = <-in:
			fmt.Println("shunt in", i)
			in = nil
			out = output
		case out <- i:
			fmt.Println("shunt out", i)
			in = input
			out = nil
		}
	}

}

func main() {
	input := make(chan int, 10)
	go func() {	// Simulate a sender to input in another part of the program
		for i := 0; i < 100; i++ {
			input <- i
		}
	}()

	acc := make(chan int, 10)	// accumulator channel
	go shunt(input, acc)
	for i := range acc {
		fmt.Println("acc read:", i)
	}
}
