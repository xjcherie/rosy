package basicsample

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}() // todo: 最后的括号是什么意思

	for i := 0; i < 2; i++ {
		// todo: 是如何知道执行msg1还是msg2的？
		select {
		case msg1 := <-c1:
			fmt.Println("receive", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}
}
