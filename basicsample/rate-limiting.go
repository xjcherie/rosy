package basicsample

import (
	"fmt"
	"time"
)

//todo: go中的时间频率限制，都需要用time自己模拟？本实例没看懂
func main() {
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests) // todo: close()在这里的作用是？

	limiter := time.Tick(200 * time.Millisecond)

	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}
	burstyLimiter := make(chan time.Time, 3) // todo: make()中第二个参数的作用？

	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	// todo: 如何判断burstyLimiter已满的？
	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()

	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)
	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}
}
