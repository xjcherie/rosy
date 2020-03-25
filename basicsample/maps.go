package basicsample

import "fmt"

// todo: basicsample()存在的作用？
func main() {
	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13
	fmt.Println("map:", m)

	v1 := m["k1"]
	fmt.Println("v1:", v1)
	fmt.Println("len:", len(m))

	delete(m, "k2")
	fmt.Println("map:", m)

	// todo: _, 是什么意思(空白标识符？)，不用该值的时候的定义？那么为什么要用两个来接收map值
	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println(n)

}
