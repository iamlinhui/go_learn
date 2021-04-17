package main

import (
	"fmt"
	"os"
	"runtime"
	"sync"
)

const BufferSize = 100

type chunk struct {
	bufSize int
	offset  int64
}

func main() {

	file, err := os.Open("/Users/lynn/Downloads/1200_2.0.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Println(err)
		return
	}

	filesize := int(fileInfo.Size())
	// Number of go routines we need to spawn.
	concurrency := filesize % BufferSize
	// buffer sizes that each of the go routine below should use. ReadAt
	// returns an error if the buffer size is larger than the bytes returned
	// from the file.
	chunkSizes := make([]chunk, concurrency)

	// All buffer sizes are the same in the normal case. Offsets depend on the
	// index. Second go routine should start at 100, for example, given our
	// buffer size of 100.
	for i := 0; i < concurrency; i++ {
		chunkSizes[i].bufSize = BufferSize
		chunkSizes[i].offset = int64(BufferSize * i)
	}

	// check for any left over bytes. Add the residual number of bytes as the
	// the last chunk size.
	if remainder := filesize % BufferSize; remainder != 0 {
		c := chunk{bufSize: remainder, offset: int64(concurrency * BufferSize)}
		concurrency++
		chunkSizes = append(chunkSizes, c)
	}

	// go1.8后 程序默认运行在多个核上面
	runtime.GOMAXPROCS(runtime.NumCPU() - 1)

	var wg sync.WaitGroup
	wg.Add(concurrency)

	for i := 0; i < concurrency; i++ {
		go func(chunkSizes []chunk, i int) {
			defer wg.Done()

			chunk := chunkSizes[i]
			buffer := make([]byte, chunk.bufSize)
			bytesRead, err := file.ReadAt(buffer, chunk.offset)

			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Printf("num %v size: %v , %v\n", i, bytesRead, string(buffer))
		}(chunkSizes, i)
	}

	wg.Wait()
}
