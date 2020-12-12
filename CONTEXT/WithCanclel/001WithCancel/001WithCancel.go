package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()
	fmt.Println("Context", ctx)
	fmt.Println("Context Error:", ctx.Err())
	fmt.Printf("Context type %T\t \n", ctx)
	fmt.Println("________")
	ctx, cancel := context.WithCancel(ctx)
	fmt.Println("ContextWithCancel", ctx)
	fmt.Println("ContextWithCancel Error:", ctx.Err())
	fmt.Printf("Context type %T\t \n", ctx)
	fmt.Println("________")

	cancel()
	fmt.Println("ContextWithCancel", ctx)
	fmt.Println("ContextWithCancel Error:", ctx.Err())
	fmt.Printf("Context type %T\t \n", ctx)

	fmt.Println("________")

}
