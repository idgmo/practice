package main

import "fmt"

func SumOfSquares(c, quit chan int) {
	y := 1
	for {
		select {
		case c <- (y * y):
			y++
		case <-quit:
			return
		}
	}
}

func main() {

	mychannel := make(chan int)

	quitchannel := make(chan int)

	sum := 0

	go func() {
		// 5 is the number of sqaures
		for i := 0; i < 5; i++ {
			sum += <-mychannel
		}
		fmt.Println(sum)
		// prints out 55
		quitchannel <- 0
	}()
	SumOfSquares(mychannel, quitchannel)
}
