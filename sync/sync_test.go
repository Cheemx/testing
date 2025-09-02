package main

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("incrementing the counter 3 times leave it at 3", func(t *testing.T) {
		counter := NewCounter()
		counter.Inc()
		counter.Inc()
		counter.Inc()

		assertCounter(t, counter, 3)
	})

	t.Run("it runs safely concurrently", func(t *testing.T) {
		wantedCount := 1000
		counter := NewCounter()

		// A WaitGroup is a counting semaphore typically used to wait for
		// a group of goroutines or tasks to finish.
		// The main goroutine calls Add to set the number of goroutines to wait for.
		// Then each of the goroutines runs and calls Done when finished.
		// At the same time, Wait can be used to block until all goroutines have finished.
		var wg sync.WaitGroup
		wg.Add(wantedCount)

		// The following will not work because
		// multiple goroutines are trying to mutate the value concurrently.
		for i := 0; i < wantedCount; i++ {
			go func() {
				counter.Inc()
				wg.Done()
			}()
		}
		wg.Wait()

		assertCounter(t, counter, wantedCount)
	})
}

func assertCounter(t testing.TB, got *Counter, want int) {
	t.Helper()
	if got.Value() != want {
		t.Errorf("got %d, want %d", got.Value(), want)
	}
}
