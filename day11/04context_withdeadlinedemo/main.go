package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	d := time.Now().Add(2000 * time.Millisecond)
	deadline, cancelFunc := context.WithDeadline(context.Background(), d)

	//尽管ctx会过期,但在任何情况下调用它的cancel函数都是很好的实践
	//如果不这样做,可能会使上下文及其父类存活的时间超过必要的时间
	defer cancelFunc()

	select {
	case <-deadline.Done():
		fmt.Println(deadline.Err())
	case <-time.After(time.Second):
		fmt.Println("overslept")
	}
}
