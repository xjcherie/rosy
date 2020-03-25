package basicsample

import (
	"fmt"
	"os"
)

//todo: go build command-line-arguments.go
// ./command-line-arguments a b c d
// 没懂这是什么用法
func main(){
	argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]

	arg := os.Args[3]

	fmt.Println(argsWithProg)
	fmt.Println(argsWithoutProg)
	fmt.Print(arg)
}