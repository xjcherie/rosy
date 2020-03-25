package basicsample

import (
	"fmt"
	"sync"
	"sync/atomic"
)

//todo:哪儿体现出来管理状态了？
func main() {
	var ops uint64

	var wg sync.WaitGroup

	for i := 0; i < 50; i++ {
		wg.Add(1) //todo: Add() 在这儿的作用是什么？

		go func() {
			for c := 0; c < 1000; c++ {
				atomic.AddUint64(&ops, 1) //todo: 为什么要传地址？，delta作用是每次加1？
			}
			wg.Done() //todo: 跟循环前的 Add()有什么关系？
		}()
	}

	wg.Wait() //todo: 作用是等待wg执行完？如何实现的？

	fmt.Println("ops:", ops)
}
