// first project main.go
package main

import (
	"fmt"
	"os"
)

func main() {

	var nameOfFile string = os.Args[1]

	file, err := os.Open(nameOfFile)
	if err != nil {
		return
	}

	defer file.Close()
	stat, err := file.Stat()
	if err != nil {
		return
	}

	var size int64 = stat.Size()

	bs := make([]byte, size)
	_, err = file.Read(bs)
	if err != nil {
		return
	}

	var count int = 0
	var itt int64 = 0
	for itt < size {
		if bs[itt] == '\n' {
			count++
		}
		itt++
	}

	fmt.Println(count)
}
