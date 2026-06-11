package main

import (
	"fmt"
	"sync"
	"time"
)

func printNumbers(label string) {
	for i := 1; i <= 5; i++ {
		fmt.Printf("%s: %d\n", label, i)
		time.Sleep(100 * time.Millisecond)
	}
}

func sendData(ch chan string, msg string) {
	time.Sleep(200 * time.Millisecond)
	ch <- msg
}

func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("worker %d mulai kerja %d\n", id, job)
		time.Sleep(300 * time.Millisecond)
		results <- job * 2
	}
}

func main() {
	fmt.Println("GOROUTINE BASIC")
	fmt.Println("---")

	go printNumbers("goroutine 1")
	go printNumbers("goroutine 2")
	time.Sleep(1 * time.Second)

	fmt.Println()
	fmt.Println("CHANNEL BASIC")
	fmt.Println("---")

	msg := make(chan string)

	go sendData(msg, "halo dari goroutine")
	received := <-msg
	fmt.Println("diterima:", received)

	fmt.Println()
	fmt.Println("BUFFERED CHANNEL")
	fmt.Println("---")

	buffered := make(chan int, 3)

	buffered <- 10
	buffered <- 20
	buffered <- 30

	fmt.Println(<-buffered)
	fmt.Println(<-buffered)
	fmt.Println(<-buffered)

	fmt.Println()
	fmt.Println("WORKER POOL")
	fmt.Println("---")

	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	for r := 1; r <= numJobs; r++ {
		<-results
	}

	fmt.Println()
	fmt.Println("SELECT")
	fmt.Println("---")

	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(200 * time.Millisecond)
		ch1 <- "dari channel 1"
	}()
	go func() {
		time.Sleep(300 * time.Millisecond)
		ch2 <- "dari channel 2"
	}()

	select {
	case msg1 := <-ch1:
		fmt.Println(msg1)
	case msg2 := <-ch2:
		fmt.Println(msg2)
	case <-time.After(500 * time.Millisecond):
		fmt.Println("timeout")
	}

	fmt.Println()
	fmt.Println("RANGE CHANNEL")
	fmt.Println("---")

	numbers := make(chan int, 3)
	numbers <- 1
	numbers <- 2
	numbers <- 3
	close(numbers)

	for n := range numbers {
		fmt.Println(n)
	}

	fmt.Println()
	fmt.Println("SYNC.WAITGROUP")
	fmt.Println("---")

	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			fmt.Printf("task %d selesai\n", id)
		}(i)
	}
	wg.Wait()
	fmt.Println("semua task selesai")
}
