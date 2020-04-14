package channel_close_test

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func dataProducer(ch chan int, wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		for i := 0; i < 100; i++ {
			ch <- i
			time.Sleep(time.Millisecond * 50)
		}
		close(ch)
		wg.Done()
	}()
}

func dataReciever(ch chan int, wg *sync.WaitGroup) {
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			for {
				if data, ok := <-ch; ok {
					fmt.Println(data)
				} else {
					break
				}
			}
			wg.Done()
		}()
	}
}

func TestChannelClose(t *testing.T) {
	wg := new(sync.WaitGroup)
	ch := make(chan int)

	dataProducer(ch, wg)
	dataReciever(ch, wg)

	wg.Wait()
}
