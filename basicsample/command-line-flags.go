package basicsample

import (
	"flag"
	"fmt"
)

func main() {
	wordPtr := flag.String("word", "foo", "a string")
	numbPtr := flag.Int("numb", 42, "an int")
	boolPtr := flag.Bool("fork", false, "a bool")

	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string var") //todo: 本行什么意思，为什么要用指针
	flag.Parse()

	fmt.Println("word:", *wordPtr) //todo: 为什么需要加指针
	fmt.Println("numb:", *numbPtr)
	fmt.Println("fork:", *boolPtr)
	fmt.Println("svar:", svar)
	fmt.Println("tail:", flag.Args())
}

// 执行过程 go build command-line-flags.go
// ./command-line-flags -word=opt -numb=7 -fork -svar=flag
// ./command-line-flags -word=opt
//todo:标志与未知参数是什么关系——标志包要求所有标志都出现在位置参数之前（否则标志将被解释为位置参数）
// ./command-line-flags -word=opt a1 a2 a3 -numb=7
// ./command-line-flags -h
// ./command-line-flags -wat