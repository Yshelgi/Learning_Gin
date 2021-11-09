package main

import "fmt"

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go func() {
		for i := 1; i <= 5; i++ {
			ch1 <- i
		}
		close(ch1)
	}()
	go func() {
		for {
			if data, ok := <-ch1; ok {
				ch2 <- data * data
			} else {
				break
			}
		}
		close(ch2)
	}()

	for i := range ch2 {
		fmt.Println(i)
	}

}
