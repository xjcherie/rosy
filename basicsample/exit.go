package basicsample

import (
	"fmt"
	"os"
)

func main() {
	defer fmt.Println("!") //todo: 为什么永远不会被执行？
	os.Exit(3) //todo: 这个3表示什么？
}
