package basicsample

import (
	"fmt"
	"time"
)

func main() {
	timer1 := time.NewTimer(2 * time.Second)

	<-timer1.C // todo: timer1.C是什么？
	fmt.Println("Timer 1 expired")

	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 expired")
	}()
	stop2 := timer2.Stop() // todo: Stop()可以直接赋值？
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}
}
