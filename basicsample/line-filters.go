package basicsample

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main(){
	scanner := bufio.NewScanner(os.Stdin)
	//todo: os.Stdin是什么？会随时监控输入的文本？
	// bufio.newScanner()是缓存的，os.Stdin()是未缓冲的？如何区分？
	// 用缓冲的包装未缓冲的，又怎么理解？

	for  scanner.Scan() { //todo: 输入一个打印一个？
		ucl := strings.ToUpper(scanner.Text())
		fmt.Println(ucl)
	}

	if err := scanner.Err(); err != nil{
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}