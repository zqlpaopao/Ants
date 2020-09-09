package main

import (
	"fmt"
	"github.com/panjf2000/ants/v2"
	"sync"
	"time"
)

func demoFunc() {
	time.Sleep(10 * time.Millisecond)
	fmt.Println("Hello World!")
}
func main() {
	/**
	###  使用之前要NewPool 不然会创建最大的协程池
	###  用time.NewTimer he sync.Cond 来定期通知回收的go程结束和清理
	###  默认回收检查是1秒
	*/
	//antss, err := ants.NewPool(1000)//初始化
	defer ants.Release()
	runTimes := 1000
	var wg sync.WaitGroup
	syncCalculateSum := func() {
		demoFunc()
		wg.Done()
	}
	for i := 0; i < runTimes; i++ {
		wg.Add(1)
		_ = ants.Submit(syncCalculateSum)
	}
	wg.Wait()
	fmt.Printf("running goroutines: %d\n", ants.Running())
	fmt.Printf("finish all tasks.\n")
}
