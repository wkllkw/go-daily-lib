package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	d := time.Now().Add(50 * time.Millisecond)
	ctx, cancel := context.WithDeadline(context.Background(), d)

	// 即使 ctx 最终会过期，作为良好的编程实践，
	// 在任何情况下都应该调用它的取消函数。
	// 如果不这样做，可能会导致 context 及其父 context
	// 比预期存活更长时间。
	defer cancel()

	select {
	case <-time.After(1 * time.Second):
		fmt.Println("ready")
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}

}
