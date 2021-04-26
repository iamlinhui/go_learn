package main

import (
	"fmt"
	"math"
	"sync"
)

func main() {

	intChan := make(chan int, 100)
	resultChan := make(chan int, 100000)

	var wg sync.WaitGroup
	wg.Add(100)

	go putNumber(intChan, 100000)

	for i := 0; i < 100; i++ {
		go calc(intChan, resultChan, &wg, i)
	}

	wg.Wait()

	close(resultChan)

	for i := range resultChan {
		fmt.Println(i)
	}

}

func calc(intChan chan int, resultChan chan int, wg *sync.WaitGroup, i int) {
	defer wg.Done()
	for {
		if num, ok := <-intChan; ok {
			fmt.Printf("管道编号:%v | num : %v,ok: %v\n", i, num, ok)
			temp := int(math.Sqrt(float64(num)))
			flag := true
			for i := 2; i <= temp; i++ {
				if num%i == 0 {
					flag = false
					break
				}
			}
			if flag {
				resultChan <- num
			}
		} else {
			break
		}
	}
}

func putNumber(intChan chan int, limit int) {
	defer close(intChan)
	for i := 1; i <= limit; i++ {
		intChan <- i
	}
}
