package main

import "fmt"

func main() {
	jobs := make(chan int, 5)
	done := make(chan interface{})

	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("received job:", j)
			} else {
				fmt.Println("received all jobs")
				done <- struct{}{}
				return
			}
		}
	}()

	for i := 1; i <= 3; i++ {
		jobs <- i
		fmt.Println("sed job", i)
	}
	close(jobs)
	fmt.Println("send all jobs")

	<-done
}
