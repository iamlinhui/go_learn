package main

import (
	"fmt"
	"sync"
	_ "time"
)

func main() {

	// 指定大小的话 是声明缓冲区  如果不指定管道大小 则需要先监听消费 并且是个注册串行的过程  只有消耗了 才能再放进去一个
	channel := make(chan int)

	fmt.Printf("%v,%v\n", len(channel), cap(channel))

	defer close(channel)

	go process(channel)

	var wg sync.WaitGroup
	wg.Add(100)

	for i := 0; i < 100; i++ {
		func(data int) {
			defer wg.Done()
			channel <- i
		}(i)
	}

	//time.Sleep(time.Second)
	wg.Wait()

}

func process(channel <-chan int) {
	//使用 select 可以解决从管道取数据的阻塞问题
	for {
		data, err := <-channel
		// 当channel关闭且数据都读完了，再读数据会读到该数据类型的零值，且第二个返回值为false
		fmt.Printf("this is %v error: %v \n", data, err)
	}

}
