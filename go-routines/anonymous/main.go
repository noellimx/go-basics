package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 10; i > 0; i-- {

		go worker(i)
		go func(_i int) {
			fmt.Printf("anon %d \n", _i)
		}(i)

	}
	time.Sleep(time.Second)
}

func worker(i int) {
	fmt.Printf("print in worker function %d \n", i)
}
