package workerpool

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func CreateTasks(count int, fn func(int) (int, error)) []Task[int, int] {
	tasks := make([]Task[int, int], count)
	for i := 0; i < 100; i++ {
		tasks[i] = Task[int, int]{
			Run:    fn,
			InArgs: 100,
		}
	}
	return tasks
}

func TestWorkerPool(t *testing.T) {

	runTask := func(input int) (int, error) {
		time.Sleep(time.Duration(input) * time.Millisecond)
		return input, nil
	}

	t.Run("simple run", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer func() { cancel() }()
		tasks := CreateTasks(100, runTask)
		pool := NewPool[int, int](ctx, 10)
		defer pool.Close()

		results, err := pool.Execute(ctx, tasks)

		require.NoError(t, err)
		require.Equal(t, 100, len(results))
	})

	t.Run("context timeout error - partial excecution", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
		defer func() { cancel() }()
		tasks := CreateTasks(100, runTask)
		pool := NewPool[int, int](ctx, 10)

		go func() {
			time.Sleep(200 * time.Millisecond)
			pool.Close()
		}()

		results, err := pool.Execute(ctx, tasks)

		require.Less(t, len(results), 100)
		require.Equal(t, err.Error(), "WorkerPool is closed")
		require.Error(t, err)
	})

	t.Run("error after closing", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
		defer func() { cancel() }()
		tasks := CreateTasks(100, runTask)
		pool := NewPool[int, int](ctx, 10)

		go func() {
			time.Sleep(100 * time.Millisecond)
			pool.Close()
		}()

		// first execution
		results, err := pool.Execute(ctx, tasks)
		require.Less(t, len(results), 100)
		require.Equal(t, err.Error(), "WorkerPool is closed")

		// second execution - afte closing
		results, err = pool.Execute(ctx, tasks)
		require.Equal(t, 0, len(results))
		require.Equal(t, err.Error(), "WorkerPool is closed")
	})

	counter := 0
	runWithError := func(input int) (int, error) {
		counter++
		if counter == 10 {
			return 0, errors.New("some error")
		}
		time.Sleep(time.Duration(input) * time.Millisecond)
		return input, nil
	}

	t.Run("error while running", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
		defer func() { cancel() }()
		tasks := CreateTasks(100, runWithError)
		pool := NewPool[int, int](ctx, 10)
		defer pool.Close()

		// first execution
		results, err := pool.Execute(ctx, tasks)
		require.Less(t, len(results), 100)
		require.Equal(t, err.Error(), "some error")
	})
}
