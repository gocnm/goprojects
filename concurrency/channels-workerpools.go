package main

import(
	"fmt"
)

func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)

	//fill up the jobs channel with 100 numbers
	for i := 0; i < 100; i++ {
		jobs <- i
	}
	//close jobs bcoz we are finished putting the stuff on to the queue/channel	
	close(jobs)

	//100 items will eventually appear on results channel
	for j := 0; j < 100 ; j ++ {
		fmt.Println(<- results)
	}

}
//worker will concurrently pulling one off at time from the jobs queue/channel and calcualte fibnaocci number
func worker(jobs <-chan int, results chan<- int) {
	for n := range jobs {
		results <- fib(n)
	}
}

func fib(n int) int {
	if n <=1 {
		return n
	}

	return fib(n-1) + fib(n-2)
}