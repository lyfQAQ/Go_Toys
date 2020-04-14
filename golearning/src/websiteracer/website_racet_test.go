package website_racer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

// func Racer(url1, url2 string) (winner string) {
// 	// startA := time.Now()
// 	// http.Get(url1)
// 	// aDuration := time.Since(startA)

// 	// startB := time.Now()
// 	// http.Get(url2)
// 	// bDuration := time.Since(startB)

// 	aDuration := measureResponseTime(url1)
// 	bDuration := measureResponseTime(url2)

// 	if aDuration < bDuration {
// 		return url1
// 	}
// 	return url2

// }

// 返回先响应的 URL。如果两个 URL 在 10 秒内都未返回结果，那么应该返回一个 error。

func Racer(a, b string) (winner string) {
	select {
	case <-ping(a):
		return a
	case <-ping(b):
		return b
	case <-time.After(5 * time.Second):
		return ""
	}
}

func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}

// 测相应时间
func measureResponseTime(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	return time.Since(start)
}

func TestRacer1(t *testing.T) {
	slowURL := "https://www.baidu.com"
	fastURL := "https://www.qq.com"

	want := fastURL
	got := Racer(slowURL, fastURL)

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestRacer2(t *testing.T) {

	// 创建服务器
	slowServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(20 * time.Millisecond)
		w.WriteHeader(http.StatusOK)
	}))

	fastServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	defer slowServer.Close()
	defer fastServer.Close()

	slowURL := slowServer.URL
	fastURL := fastServer.URL

	want := fastURL
	got := Racer(slowURL, fastURL)

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}

}
