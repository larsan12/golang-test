package main

import (
	"context"
	"fmt"
	"math/rand"
	"route256/libs/workerpool"
	"time"
)

func run(input int) int {
	time.Sleep(time.Duration(input) * time.Millisecond)
	return input
}

func main() {

	count := 100
	tasks := make([]workerpool.Task[int, int], count)

	for i := 0; i < 100; i++ {
		tasks[i] = workerpool.Task[int, int]{
			Run:    run,
			InArgs: rand.Intn(1000),
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer func() { cancel() }()

	pool := workerpool.NewPool[int, int](ctx, 10)

	go func() {
		time.Sleep(2 * time.Second)
		pool.Close()
	}()

	results, err := pool.Execute(ctx, tasks)
	if err != nil {
		fmt.Println("error ", err)
	}
	fmt.Println("results len", len(results))
	fmt.Println("results ", results)

	results, err = pool.Execute(ctx, tasks)
	if err != nil {
		fmt.Println("2error ", err)
	}
	fmt.Println("2results len", len(results))
	fmt.Println("2results ", results)
}
