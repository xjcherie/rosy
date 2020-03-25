package basicsample

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println
	t := time.Now()
	p(t.Format(time.RFC3339)) //todo:time.RFC3339等名称来源

	t1, e := time.Parse(time.RFC3339, "2019-11-03T19:56:05+00:00")
	p(t1)

	p(t.Format("3:04PM"))
	p(t.Format("Mon Jan _2 15:04:05 2016"))
	p(t.Format("2016-01-02T15:04:05.999999-07:00"))
	form := "3 04 PM"
	t2, e := time.Parse(form, "8 41 PM")
	p(t2) //todo:0000-01-01 20:41:00 +0000 UTC，为什么是这样的输出结果？

	fmt.Printf("%d-%02d-%02dT%02d:%02d:%02d-00\n", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
	ansic := "Mon Jan _2 15:04:05 2016"
	_, e = time.Parse(ansic, "8:41PM")
	p(e)
}
