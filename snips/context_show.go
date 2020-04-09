package main

import (
	"context"
	"fmt"
)

func main() {
	fmt.Println("Hello, playground")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	ctx = context.WithValue(ctx, "int", 2)
	ctx = context.WithValue(ctx, "string", "got it")

	fmt.Println(ctx.Value("int"), ctx.Value("string"))

}
