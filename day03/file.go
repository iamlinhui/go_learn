package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {

	var data = make([]byte, 256)
	//os.File.Read(data)

	file, err := os.Open("/Users/lynn/Downloads/1200_2.0.txt")

	if err != nil {
		panic(err)
	}

	defer file.Close()

	reader := bufio.NewReader(file)

	for {
		index, err := reader.Read(data)
		if err != nil {
			if err != io.EOF {
				fmt.Println(err)
			}
			break
		}
		fmt.Println(string(data[:index]))
	}

}
