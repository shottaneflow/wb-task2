package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	var countWorkers int
	fmt.Fscan(os.Stdin, &countWorkers)
	channel := make(chan int, 1)
	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, syscall.SIGINT)
	for i := 0; i < countWorkers; i++ {
		go func(numWorker int) {
			for n := range channel {
				fmt.Printf("Worker %d takes value from channel: %d\n", numWorker, n)
			}
		}(i)
	}
	gracefulShutdown := func() {
		fmt.Println("Завершение работы программы")
		close(channel)
		close(stopChan)
	}
	for {
		select {
		case <-stopChan:
			gracefulShutdown()
			os.Exit((0))
		default:
			channel <- rand.Int()
		}

	}
}
