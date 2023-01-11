package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

var GLOBAL_SEPHAMORE_1 sync.WaitGroup

func BLOCK() {
	GLOBAL_SEPHAMORE_1.Add(1)

}
func BARRIER() {
	GLOBAL_SEPHAMORE_1.Wait()
}

func LOGCOUNT() {

}

func Synchronize() {

	GLOBAL_SEPHAMORE_1.Add(1)
	go task1()
	BARRIER()

	GLOBAL_SEPHAMORE_1.Add(1)

	go task2()
	BARRIER()

	GLOBAL_SEPHAMORE_1.Add(1)
	go task3()
	BARRIER()

	GLOBAL_SEPHAMORE_1.Add(1)
	go task4()
	BARRIER()

	log.Println("LAST BARRIER")
}

func main() {

}

func task1() {
	fmt.Println("go task 1")
	time.Sleep(30 * time.Millisecond)
	fmt.Println("task 1")
	GLOBAL_SEPHAMORE_1.Done()
}

func task2() {
	fmt.Println("go task 2")

	time.Sleep(50 * time.Millisecond)
	fmt.Println("task 2")
	GLOBAL_SEPHAMORE_1.Done()
}

func task3() {
	fmt.Println("go task 3")

	time.Sleep(10 * time.Millisecond)
	fmt.Println("task 3")
	GLOBAL_SEPHAMORE_1.Done()
}

func task4() {
	fmt.Println("go task 4")

	time.Sleep(20 * time.Millisecond)
	fmt.Println("task 4")
	GLOBAL_SEPHAMORE_1.Done()
}
