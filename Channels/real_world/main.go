package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs chan int, results chan int) {
	for j := range jobs {
		fmt.Println("Worker", id, "processing job", j)
		time.Sleep(time.Second)
		results <- j * 2
	}
}

func main() {
	jobs := make(chan int, 5)
	results := make(chan int, 5)

	//start workers
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	//send jobs
	for j := 1; j <= 5; j++ {
		jobs <- j
	}

	close(jobs)

	for a := 1; a <= 5; a++ {
		fmt.Println(<-results)
	}
}
