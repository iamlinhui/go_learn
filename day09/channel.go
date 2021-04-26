package main

import (
	"fmt"
	"math"
	"sync"
)

func main() {

	intChan := make(chan int, 100)
	resultChan := make(chan int, 100)

	var wg sync.WaitGroup
	wg.Add(100)

	go putNum(intChan, 100000)

	for i := 0; i < 100; i++ {
		go calc(intChan, resultChan, &wg)
	}

	wg.Wait()

	close(resultChan)

	for i := range resultChan {
		fmt.Println(i)
	}

}

func calc(intChan chan int, resultChan chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		if num, ok := <-intChan; ok {
			fmt.Printf("num : %v,ok: %v\n", num, ok)
			flag := int(math.Sqrt(float64(num)))
			for i := 2; i <= flag; i++ {
				if num%i == 0 {
					resultChan <- num
				}
			}
		} else {
			break
		}
	}
}

func putNum(intChan chan int, limit int) {
	defer close(intChan)
	for i := 1; i <= limit; i++ {
		intChan <- i
	}
}
