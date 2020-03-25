package basicsample

import (
	"fmt"
	"sort"
)

type byLength [] string //todo: 在这个位置定义的作用？

func (s byLength) Len() int {
	return len(s)
}

func (s byLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s byLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

// todo:是怎么使用的Swap和Less函数？自动替换sort内部函数？如何实现的
func main() {
	fruits := [] string{"peach", "banana", "kiwi"}
	sort.Sort(byLength(fruits))
	fmt.Println(fruits)
}
