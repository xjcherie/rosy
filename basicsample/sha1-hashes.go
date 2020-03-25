package basicsample

import (
	"crypto/sha1"
	"fmt"
)

// todo: go中提供了多少种加解密方法，其中包含集中hash函数
// todo: 哈希强度？跟什么有关，如何增强
func main() {
	s := "sha1 this string"
	h := sha1.New()

	h.Write([]byte(s))
	bs := h.Sum(nil) //todo: nil在这里的作用，添加了其它字符串后，好像会有一定规律

	fmt.Println(s)
	fmt.Printf("%x\n", bs) //todo: 以16进制打印，%x是16进制，其它呢
}
