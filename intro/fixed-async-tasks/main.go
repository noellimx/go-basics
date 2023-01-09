package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	done := make(chan string)
	now := time.Now()
	go task1(done)
	go task2(done)
	go task3(done)
	go task4(done)

	log.Println(<-done)
	log.Println(<-done)
	log.Println(<-done)
	log.Println(<-done)

	fmt.Println("elapsed:", time.Since(now))
}

func task1(done chan string) {
	time.Sleep(100 * time.Millisecond)
	done <- "task1"
}

func task2(done chan string) {
	time.Sleep(200 * time.Millisecond)
	done <- "task2"
}

func task3(done chan string) {
	done <- "task3"
}

func task4(done chan string) {
	time.Sleep(100 * time.Millisecond)
	done <- "task4"
}
