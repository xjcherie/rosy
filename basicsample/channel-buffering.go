package basicsample

import "fmt"

func main() {
	// todo: go中有几种数据类型
	messages := make(chan string, 2)

	messages <- "buffered"
	messages <- "channel"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
