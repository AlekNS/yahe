package utils

import (
	"context"
	"errors"
	"testing"
	"time"
)

func TestRepeatableMethodSuccess(t *testing.T) {
	err := RepeatMethod(context.TODO(), BackoffEmpty(), 3, func(context.Context) error {
		return nil
	})
	if err != nil {
		t.Fatalf("Expected nil, got %v", err)
	}
}

func TestRepeatableMethodFail(t *testing.T) {
	err := RepeatMethod(context.TODO(), BackoffLinear(time.Millisecond), 3, func(context.Context) error {
		return errors.New("error")
	})
	if err == nil {
		t.Fatalf("Expected error, got nil")
	}
}

func TestRepeatableMethodStopOnContextDone(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	err := RepeatMethod(ctx, BackoffLinear(time.Second), 3, func(context.Context) error {
		return errors.New("error")
	})
	if err != context.Canceled {
		t.Fatal("Expected error to be context.Canceled, got", err)
	}
}

func TestRepeatableMethodSuccessOnSecondTime(t *testing.T) {
	times := 4
	err := RepeatMethod(context.TODO(), BackoffExponential(time.Millisecond, 2), 3, func(context.Context) error {
		times--
		if times > 0 {
			return errors.New("error")
		}
		return nil
	})
	if err != nil {
		t.Fatalf("Expected nil, got %v", err)
	}
	if times != 0 {
		t.Fatalf("Expected times == 0, got %v", times)
	}
}

func TestRepeatableMethodFailForPanics(t *testing.T) {
	err := RepeatMethod(context.TODO(), BackoffEmpty(), 3, WrapPanic(func(context.Context) {
		panic("panic")
	}))
	if err == nil || err.Error() != "panic" {
		t.Fatalf("Expected error with panic, got %v", err)
	}
}
