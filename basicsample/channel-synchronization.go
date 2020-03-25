package basicsample

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true
}

func main() {
	done := make(chan bool, 1)
	go worker(done)

	// todo:为什么有这句会等待执行，没有就不等函数调用完，通道会一直等待有值才执行该句？
	<-done
}
