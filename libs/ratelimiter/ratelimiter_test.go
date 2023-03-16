package ratelimiter

import (
	"context"
	"math"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestRateLimiter(t *testing.T) {

	// generic test
	testRequest := func(count int, rps uint, burst uint) {
		ctx := context.Background()
		limiter := NewLimiter(rps, burst)
		defer limiter.Close()

		// считаем время выполнения
		t1 := time.Now()
		for i := 0; i < count; i++ {
			limiter.Wait(ctx)
		}
		t2 := time.Now()

		// округляем секунды
		diff := int(math.Round(t2.Sub(t1).Seconds()))

		// считаем теоритическое время выполнения
		expectedDiff := int(
			math.Round(
				float64(count-int(burst)) / float64(rps),
			),
		)

		require.Equal(t, expectedDiff, diff)
	}

	t.Run("simple test", func(t *testing.T) {
		testRequest(100, 20, 0)
	})

	t.Run("simple test", func(t *testing.T) {
		testRequest(100, 50, 10)
	})

	t.Run("simple test", func(t *testing.T) {
		testRequest(100, 20, 20)
	})

	t.Run("simple test", func(t *testing.T) {
		testRequest(100, 10, 10)
	})
}
