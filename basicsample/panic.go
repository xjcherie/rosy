package basicsample

import "os"

func main() {
	panic("a problem") //todo: 这句作用是什么，跟后边的panic(err)什么关系

	_, err := os.Create("/tmp/file")
	if err != nil {
		panic(err)
	}
}
