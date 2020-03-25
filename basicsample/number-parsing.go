package basicsample

import (
	"fmt"
	"strconv"
)

func main() {
	f, _ := strconv.ParseFloat("1.234", 64)
	fmt.Println(f)

	i, _ := strconv.ParseInt("123", 0, 64) //todo: 0与64的作用
	fmt.Println(i)

	d, _ := strconv.ParseInt("0x1c8", 0, 64)
	fmt.Println(d)

	u, _ := strconv.ParseUint("789", 0, 64) //todo: uint与int的区别
	fmt.Println(u)

	k, _ := strconv.Atoi("135") //todo: Atoi是什么？
	fmt.Println(k)

	_, e := strconv.Atoi("wat")
	fmt.Println(e)
}
