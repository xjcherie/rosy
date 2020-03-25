package basicsample

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	fooCmd := flag.NewFlagSet("foo", flag.ExitOnError) //todo: flag是什么？
	fooEnable := fooCmd.Bool("enable", false, "enable")
	fooName := fooCmd.String("name", "", "name")

	barCmd := flag.NewFlagSet("bar", flag.ExitOnError)
	barLevel := barCmd.Int("level", 0, "level")

	//todo: 子命令应作为程序的第一个参数，不懂
	if len(os.Args) < 2 {
		fmt.Println("expected 'foo' or 'bar' subcommands")
		os.Exit(1) //todo: 1表示什么，可以有几种输入
	}

	switch os.Args[1] {
	case "foo":
		fooCmd.Parse(os.Args[2:])
		fmt.Println("subcommand 'foo'")
		fmt.Println(" enable:", *fooEnable)
		fmt.Println(" name:", *fooName)
		fmt.Println(" tail:", fooCmd.Args())
	case "bar":
		barCmd.Parse(os.Args[2:])
		fmt.Println("subcommand 'bar'")
		fmt.Println(" level:", *barLevel)
		fmt.Println(" tail:", barCmd.Args())
	default:
		fmt.Println("expected 'foo' or 'bar' subcommands")
		os.Exit(1)
	}
}

// todo: 不太懂输出结果，第一个没有输入enable的值，为什么输出是true？第二个"-level 8"没有等于号，为什么level输出的是8
// go build command-line-subcommands.go
// 1. ./command-line-subcommands foo -enable -name=joe a1 a2
// 2. ./command-line-subcommands bar -level 8 a1
// 3. ./command-line-subcommands bar -enable a1
