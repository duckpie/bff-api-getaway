package graph

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Effector func(context.Context) (interface{}, error)

func Repeater(effector Effector, retries int, delay time.Duration) Effector {
	return func(ctx context.Context) (interface{}, error) {
		for r := 0; ; r++ {
			resp, err := effector(ctx)
			if err == nil {
				s, _ := status.FromError(err)
				if s.Code() != codes.Unavailable || r >= retries {
					return resp, err
				}
			}

			delay += time.Second
			fmt.Printf("Attempt %d failed; retrying in %v\n", r+1, delay)

			select {
			case <-time.After(delay):
			case <-ctx.Done():
				return nil, ctx.Err()
			}
		}
	}
}
