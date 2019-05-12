package utils

import (
	"context"
	"fmt"
	"time"
)

// WrapPanic wraps any method for recovering after panic
func WrapPanic(method func(ctx context.Context)) func(ctx context.Context) error {
	return func(ctx context.Context) (err error) {
		defer func() {
			if e := recover(); e != nil {
				err = fmt.Errorf("%+v", e)
			}
		}()
		method(ctx)
		return err
	}
}

// BackoffFunc method
type BackoffFunc func(ctx context.Context) error

// BackoffEmpty no waits
func BackoffEmpty() BackoffFunc {
	return func(ctx context.Context) error {
		return nil
	}
}

// BackoffLinear waits linear time
func BackoffLinear(waitInterval time.Duration) BackoffFunc {
	return func(ctx context.Context) error {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(waitInterval):
			return nil
		}
	}
}

// BackoffExponential waits exponential time
func BackoffExponential(waitInterval time.Duration, multiply float64) BackoffFunc {
	interval := waitInterval
	return func(ctx context.Context) error {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(interval):
			interval = time.Duration(float64(interval) * multiply)
			return nil
		}
	}
}

// RepeatMethod for max attempts and wait interval before them in sync mode
func RepeatMethod(ctx context.Context, backoff BackoffFunc, maxAttempts int, callback func(context.Context) error) (err error) {
	for {
		err = callback(ctx)
		if err == nil {
			return nil
		}
		if ctx.Err() != nil {
			return ctx.Err()
		}

		maxAttempts--
		if maxAttempts < 0 {
			break
		}

		err = backoff(ctx)
		if err != nil {
			return err
		}
	}
	return err
}
