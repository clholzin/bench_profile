package main

import (
	"context"
	"fmt"
	"html"
	"log"
	"net/http"
	_ "net/http/pprof"
	"runtime/trace"
	"time"
)

var fooey func(w http.ResponseWriter, r *http.Request)

var st map[string][]byte = make(map[string][]byte)

func main() {
	log.Println("Starting Server")

	Start()
}

func Start() {
	stats := NewStat()
	http.HandleFunc("/", stats.Foo())
	log.Fatal(http.ListenAndServe(":8080", nil))
}

type httpStats struct {
	mem      map[string]int
	reqcount int64
}

func NewStat() httpStats {
	h := httpStats{make(map[string]int), 0}
	return h
}

func (h httpStats) Foo() func(w http.ResponseWriter, r *http.Request) {
	fooey = func(w http.ResponseWriter, r *http.Request) {

		ctx, task := trace.NewTask(context.Background(), "Foo http route")
		defer task.End()

		path := html.EscapeString(r.URL.Path)

		// Log with context and uuid
		trace.Log(ctx, "Foo", "Execute Foo route work flow")

		h.mem[path]++
		fmt.Printf("got request # %d\n", h.mem[path])
		reqNumber := h.reqcount + 1
		h.reqcount = reqNumber
		trace.WithRegion(ctx, "Foo do work", work(ctx, reqNumber))
		fmt.Fprintf(w, "Profiling is awesome sauce, %q", html.EscapeString(r.URL.Path))
	}
	return fooey

}

func work(ctx context.Context, n int64) func() {
	return func() {
		trace.Log(ctx, "Execute Foo route work flow", "")
		//time.Sleep((time.Millisecond * 10) * time.Duration(n))

		for i := int64(0); i < n; i++ {
			t := time.Now().String()
			st[t] = make([]byte, 512)
		}

		Fib(int(n))

		// fmt.Println("finishing request ", n)
	}
}

func Fib(n int) int {
	switch n {
	case 0:
		return 0
	case 1:
		return 1
	case 2:
		return 1
	default:
		return Fib(n-1) + Fib(n-2)
	}
}
