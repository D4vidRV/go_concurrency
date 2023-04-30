package main

import "fmt"

func main() {
	// input
	nums := []int{2, 3, 4, 5, 1, 7}
	// stage 1
	dataChannel := sliceToChannel(nums) // the channel out  of sliceToChannel
	// stage2
	finalChannel := sq(dataChannel) //is the in of sq function, so both are coordinated
	// stage 3
	for v := range finalChannel {
		fmt.Println(v)
	}
}

func sliceToChannel(nums []int) <-chan int {
	out := make(chan int)

	go func() {
		for _, num := range nums {
			out <- num
		}
		close(out)
	}()
	return out
}

func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}
