package onelabhomework5

import (
	"context"
	"fmt"
	"testing"
)

func withError1() error {
	return fmt.Errorf("some error")
}

func withoutError1() error {
	return nil
}
func TestExecute1(t *testing.T) {
	numberOfErrFuncs := 10
	numberOfFuncs := 25

	tasks := make([]func() error, 0, 25)

	for i := 0; i < numberOfFuncs; i++ {
		if i < numberOfErrFuncs {
			tasks = append(tasks, withError1)
		} else {
			tasks = append(tasks, withoutError1)
		}
	}

	for i := 0; i < numberOfFuncs; i++ {
		if i <= numberOfErrFuncs {
			if err := Execute1(tasks, i); err == nil {
				t.Errorf("TestExecute2 => got (%v), want (%v)", nil, ErrErrLimitExceeded)
			}
		} else {
			if err := Execute1(tasks, i); err != nil {
				t.Errorf("TestExecute2 => got (%v), want (%v)", err, nil)
			}
		}
	}
}

///////////////////////////////////////////////////////////////////////////////////////////////

func withError2(ctx context.Context) error {
	select {
	case <-ctx.Done():
		return ctx.Err()

	default:
		return fmt.Errorf("errrr")
	}
}

func withoutError2(ctx context.Context) error {
	select {
	case <-ctx.Done():
		return ctx.Err()

	default:
		return nil
	}
}

func TestExecute2(t *testing.T) {
	numberOfErrFuncs := 10
	numberOfFuncs := 25

	tasks := make([]func(context.Context) error, 0, 25)

	for i := 0; i < numberOfFuncs; i++ {
		if i < numberOfErrFuncs {
			tasks = append(tasks, withError2)
		} else {
			tasks = append(tasks, withoutError2)
		}
	}

	for i := 0; i < numberOfFuncs; i++ {
		if i <= numberOfErrFuncs {
			if err := Execute2(tasks, i); err == nil {
				t.Errorf("TestExecute2 => got (%v), want (%v)", nil, ErrErrLimitExceeded)
			}
		} else {
			if err := Execute2(tasks, i); err != nil {
				t.Errorf("TestExecute2 => got (%v), want (%v)", err, nil)
			}
		}
	}
}
