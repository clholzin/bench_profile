package main

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"strings"
	"testing"
	"time"
)

func init() {
	//go Start()
	time.Sleep(time.Second * 3)

}

func TestMain(m *testing.M) {
	// setup before tests
	//go main()
	os.Exit(m.Run())
	// teardown after tests
}

func TestFoo(t *testing.T) {
	httpStat := NewStat()
	req := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	fooHandler := httpStat.Foo()
	fooHandler(w, req)
	resp := w.Result()
	expected := 200
	if resp.StatusCode != expected {
		t.Errorf("Got %d want %d", resp.StatusCode, expected)
	}
}

func TestWork(t *testing.T) {
	var worker = work(context.Background(), 20)
	t.Log(reflect.TypeOf(worker).Kind().String())
	if reflect.TypeOf(worker).Kind().String() != "func" {
		t.Logf("worker was not type func")
	}
}

func BenchmarkFoo(b *testing.B) {
	//go Start()
	time.Sleep(time.Second * 3)
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
