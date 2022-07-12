package ctx

import (
	"context"
	"fmt"
	"time"
)

func HandleRequest(ctx context.Context) {
	go WriteRedis(ctx)
	go WriteDatabase(ctx)

	for {
		select {
		case <-ctx.Done():
			fmt.Println("HandleRequest Done.")
			return
		default:
			fmt.Println("HandleRequest Running...")
			time.Sleep(2 * time.Second)
		}
	}
}

func WriteRedis(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("WriteRedis Done.")
			return
		default:
			fmt.Println("WriteRedis Running...")
			time.Sleep(2 * time.Second)
		}
	}
}

func WriteDatabase(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("WriteDatabase Done.")
			return
		default:
			fmt.Println("WriteDatabase Running...")
			time.Sleep(2 * time.Second)
		}
	}
}
