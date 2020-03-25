package basicsample

import (
	"fmt"
	s "strings"
)

var p = fmt.Println // todo: 函数直接赋值给变量？程序中为什么会出现"未处理的错误"？
// todo: 有多少种常用的编码方式 utf-8……

func main() {
	p("Contains:", s.Contains("test", "es"))
	p("Count", s.Count("test", "t"))
	p("HasPrefix:", s.HasPrefix("test", "te"))
	p("HasSuffix:", s.HasSuffix("test", "st"))
	p("Index:", s.Index("test", "e"))
	p("Join:", s.Join([]string{"a", "b"}, "-"))
	p("Repeat:", s.Repeat("a", 5))
	p("Replace:", s.Replace("foo", "o", "0", -1))
	p("Replace:", s.Replace("fooo", "o", "0", 1))
	p("Split:", s.Split("a-b-c-d-e", "-"))
	p("ToLower:", s.ToLower("TEST"))
	p("ToUpper:", s.ToUpper("test"))
	p()
	p("Len:", len("hello"))
	p("Char:", "hello"[1])
}