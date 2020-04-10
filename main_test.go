package main

import (
	"context"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
	"time"
)

func init() {
	go Start()
	time.Sleep(time.Second * 3)

}

func BenchmarkFoo(b *testing.B) {

	count := 0

	for i := 0; i < b.N; i++ {
	retry:
		resp, err := http.Get("http://localhost:8080/")
		if err != nil {
			count++
			if count > 5 {
				b.FailNow()
				return
			}
			goto retry
			// handle error
		}
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)
		if strings.Contains(string(body), "Profiling is awesome sauce") {
			// great
		} else {
			b.FailNow()
		}

	}
}

func BenchmarkWork(b *testing.B) {
	for n := 0; n < b.N; n++ {
		work(context.Background(), 20)()
	}
}

var Result int64

func BenchmarkFib20(b *testing.B) {
	var d int
	for n := 0; n < b.N; n++ {
		d += Fib(20) // run the Fib function b.N times
	}
	Result = int64(d) // prevent compiler optimizing
}

func BenchmarkFib28(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Fib(28) // run the Fib function b.N times
	}
}
