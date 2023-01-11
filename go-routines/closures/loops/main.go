package main

import (
	"fmt"
	"time"
)

type request func()

func main() {

	variablePerIteration()
	variablePerLoop()
}

func variablePerLoop() {

	requests := make([]request, 0)
	for i := 1; i <= 100; i++ {
		f := func() {
			time.Sleep(200 * time.Millisecond)
			fmt.Println("[variablePerLoop] request", i)
		}
		requests = append(requests, f)
	}

	for _, r := range requests {
		r()
	}

}

func variablePerIteration() {

	requests := make([]request, 0)
	for i := 1; i <= 100; i++ {
		i := i // shadow
		f := func() {
			time.Sleep(200 * time.Millisecond)
			fmt.Println("[variablePerIteration] request", i)
		}
		requests = append(requests, f)
	}

	for _, r := range requests {
		r()
	}
}
