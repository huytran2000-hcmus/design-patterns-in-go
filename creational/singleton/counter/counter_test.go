package counter

import (
	"sync"
	"testing"
)

func TestGetInstance(t *testing.T) {
	t.Run("synchronously", testGetInstance)
	t.Run("asynchronously", testAsyncGetInstance)
}

func TestCounter_Incr(t *testing.T) {
	t.Run("synchronously", testCounterIncr)
	t.Run("asynchronously", testAsyncCounterIncr)
}

func testGetInstance(t *testing.T) {
	want := GetInstance()
	got := GetInstance()

	if got != want {
		t.Errorf("GetInstance() = %v, want %v", got, want)
	}
}

func testAsyncGetInstance(t *testing.T) {
	for i := 0; i < 100; i++ {
		singleton = nil
		wantCh := asyncGetInstance()
		gotCh := asyncGetInstance()

		want := <-wantCh
		got := <-gotCh

		if got != want {
			t.Errorf("GetInstance() = %v, want %v", got, want)
		}
	}
}

func asyncGetInstance() chan *Counter {
	wantCh := make(chan *Counter)
	go func() {
		wantCh <- GetInstance()
	}()
	return wantCh
}

func testCounterIncr(t *testing.T) {
	t.Parallel()
	singleton = nil
	counter := GetInstance()

	count := 100
	for i := 0; i < count; i++ {
		counter.Incr()
	}

	want := count
	got := counter.Get()

	if got != want {
		t.Errorf("call *counter.Incr() %d times, *counter.Get() = %d, want %d", count, got, want)
	}
}

func testAsyncCounterIncr(t *testing.T) {
	t.Parallel()
	singleton = nil
	counter := GetInstance()
	count := 100
	asyncIncr(counter, count)

	want := count
	got := counter.Get()

	if got != want {
		t.Errorf("call *counter.Incr() %d times, *counter.Get() = %d, want %d", count, got, want)
	}
}

func asyncIncr(counter *Counter, n int) {
	var wg sync.WaitGroup

	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			counter.Incr()
			wg.Done()
		}()
	}

	wg.Wait()
}
