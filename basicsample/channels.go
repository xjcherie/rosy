package basicsample

import "fmt"

//todo: 还是不太懂通道的意思
func main() {
	messages := make(chan string)
	// todo: 这个怎么算有接收者
	go func() { messages <- "ping" }()

	msg := <-messages
	fmt.Println(msg)
}
