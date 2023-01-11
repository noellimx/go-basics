package main

import (
	"log"
	"time"
)

func main() {
	t1 := time.Now()
	for i := 1; i < 5; i++ {
		ms := time.Duration(i) * 100
		immediateTask(ms)
	}
	time.Sleep(6)

	t2 := time.Now()
	log.Printf("time elapsed%s", t2.Sub(t1))

}

func immediateTask(ms time.Duration) {

	time.Sleep(ms)
	log.Printf("%v", ms)
}
