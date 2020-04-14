package share_mem

import (
	"sync"
	"testing"
	"time"
)

// 互斥量
func TestCounterWithMutex(t *testing.T) {
	var mut sync.Mutex

	counter := 0

	for i := 0; i < 5000; i++ {
		go func() {
			defer func() {
				mut.Unlock()
			}()
			mut.Lock()
			counter++
		}()
	}
	time.Sleep(1 * time.Second)
	t.Logf("Counter = %d", counter)
}

func TestCounterWaitGroup(t *testing.T) {
	var wg sync.WaitGroup
	var mut sync.Mutex
	counter := 0

	for i := 0; i < 5000; i++ {
		wg.Add(1)
		go func() {
			defer func() {
				mut.Unlock()
			}()
			mut.Lock()
			counter++
			wg.Done()
		}()
	}
	wg.Wait()
	t.Logf("Counter = %d", counter)
}
