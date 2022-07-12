package ctx

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestHandleRequest(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	go HandleRequest(ctx)

	time.Sleep(5 * time.Second)
	fmt.Println("It's time to stop all sub goroutines!")

	cancel()

	time.Sleep(5 * time.Second)
}
