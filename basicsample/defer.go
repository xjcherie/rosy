package basicsample

import (
	"fmt"
	"os"
)

// todo: 为什么看不到被创建的文件？创建的是什么文件
func createFile(p string) *os.File {
	fmt.Println("creating")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(f *os.File) {
	fmt.Println("writing")
	fmt.Fprintln(f, "data") //todo: fmt.Fprintln用于写文件？
	fmt.Fprintln(f, "haha")
}

func closeFile(f *os.File) {
	fmt.Println("closing")
	err := f.Close()
	if err != nil {
		fmt.Fprintln(os.Stderr, "error: %v\n", err) //todo:报错提示是什么？为什么并没有出现错误
		os.Exit(1)
	}
}

func main() {
	f := createFile("./defer.txt")
	defer closeFile(f)
	writeFile(f)
}
