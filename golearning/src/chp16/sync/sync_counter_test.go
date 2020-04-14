package sync

import (
	"sync"
	"testing"
)

type Counter struct {
	value int
	mu    sync.Mutex
}

func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func (c *Counter) Value() int {
	return c.value
}

/* 一个 WaitGroup 等待一组 goroutine 的完成。
主 goroutine 调用 Add 来设置等待的 goroutine 的数量。
然后每一个 goroutines 运行并在完成时调用 Done。
同时，可以使用 Wait 来阻塞，直到所有的 goroutines 完成。*/

func TestCounter(t *testing.T) {
	t.Run("incrementing the counter 3 times leaves it at 3", func(t *testing.T) {
		wantedCount := 1000
		counter := Counter{}

		var wg sync.WaitGroup
		wg.Add(wantedCount)

		for i := 0; i < wantedCount; i++ {
			go func(w *sync.WaitGroup) {
				counter.Inc()
				w.Done()
			}(&wg)
		}
		wg.Wait()

		if counter.Value() != wantedCount {
			t.Errorf("got %d, want %d", counter.Value(), wantedCount)
		}
	})
}
