package main

import (
	"log"
	"time"
)

type STRUCT = string

func main() {
	done := make(chan STRUCT)
	go task1(done)
	log.Printf(<-done)
	go task2(done)
	log.Printf(<-done)
	go task3(done)
	log.Printf(<-done)

}

func task1(done chan STRUCT) {
	time.Sleep(100 * time.Millisecond)
	done <- "task1 - 100ms"
}

func task2(done chan STRUCT) {
	time.Sleep(50 * time.Millisecond)
	done <- "task2 - 50"
}

func task3(done chan STRUCT) {
	time.Sleep(10 * time.Millisecond)
	done <- "task3 - 10ms"
}
