package main

import (
	"log"
	"os"
	"sync"
	"time"
)

var LoggerTo1 *log.Logger
var LoggerTo2 *log.Logger

func init() {

	f, err := os.OpenFile("text1.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	f2, err := os.OpenFile("text2.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f2.Close()
	LoggerTo1 = log.New(f, "prefix", log.LstdFlags)
	LoggerTo2 = log.New(f, "prefix", log.LstdFlags)
}
func main() {
	namedClosures()
	anonymousClosures()
}

var cycles = 1000000

func namedClosures() {
	var wg sync.WaitGroup
	f := func(n int, _wg *sync.WaitGroup) {
		defer _wg.Done()
		LoggerTo1.Println("printing from named closure", n)
	}

	t1 := time.Now()
	for i := 0; i < cycles; i++ {
		wg.Add(1)
		go f(i, &wg)
	}

	t2 := time.Now()
	log.Printf("[namedClosures] time elapsed %d", t2.Sub(t1))
	wg.Wait()
}

func anonymousClosures() {
	var wg sync.WaitGroup
	t1 := time.Now()

	for i := 0; i < cycles; i++ {
		wg.Add(1)
		go func(n int, _wg *sync.WaitGroup) {
			defer _wg.Done()
			LoggerTo2.Println("printing from anonymoys closure", n)
		}(i, &wg)
	}
	t2 := time.Now()
	log.Printf("[anonymousClosures] time elapsed %d", t2.Sub(t1))
	wg.Wait()
}
