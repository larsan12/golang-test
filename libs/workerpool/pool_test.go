package workerpool

import (
	"context"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func RunTask(input int) int {
	time.Sleep(time.Duration(input) * time.Millisecond)
	return input
}

func CreateTasks(count int) []Task[int, int] {
	tasks := make([]Task[int, int], count)
	for i := 0; i < 100; i++ {
		tasks[i] = Task[int, int]{
			Run:    RunTask,
			InArgs: rand.Intn(1000),
		}
	}
	return tasks
}

func TestWorkerPool(t *testing.T) {

	t.Run("simple run", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
		defer func() { cancel() }()
		tasks := CreateTasks(100)
		pool := NewPool[int, int](ctx, 10)

		results, err := pool.Execute(ctx, tasks)

		require.Equal(t, 100, len(results))
		require.NoError(t, err)
	})

	t.Run("context timeout error - partial excecution", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
		defer func() { cancel() }()
		tasks := CreateTasks(100)
		pool := NewPool[int, int](ctx, 10)

		go func() {
			time.Sleep(2 * time.Second)
			pool.Close()
		}()

		results, err := pool.Execute(ctx, tasks)

		require.Greater(t, 100, len(results))
		require.Equal(t, err.Error(), "WorkerPool is closed")
		require.Error(t, err)
	})

	t.Run("error after closing", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
		defer func() { cancel() }()
		tasks := CreateTasks(100)
		pool := NewPool[int, int](ctx, 10)

		go func() {
			time.Sleep(time.Second)
			pool.Close()
		}()

		// first execution
		results, err := pool.Execute(ctx, tasks)
		require.Greater(t, 100, len(results))
		require.Equal(t, err.Error(), "WorkerPool is closed")

		// second execution - afte closing
		results, err = pool.Execute(ctx, tasks)
		require.Equal(t, 0, len(results))
		require.Equal(t, err.Error(), "WorkerPool is closed")
	})
}
