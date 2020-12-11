package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()
	fmt.Println("context:", ctx)
	fmt.Println("context error:", ctx.Err())
	fmt.Printf("type %T:", ctx)

}
