package onelabhomework5

import (
	"context"
	"errors"
	"sync"
)

func Execute2(tasks []func(ctx context.Context) error, E int) error {
	wg := sync.WaitGroup{}
	mx := sync.Mutex{}

	errCount := 0

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for _, f := range tasks {
		wg.Add(1)

		go func(f func(context.Context) error, ctx context.Context) {
			err := f(ctx)

			if err != nil && !errors.Is(err, ctx.Err()) {
				mx.Lock()

				if errCount < E {
					// fmt.Println("error ctx", err)
					errCount++
				} else {
					cancel()
				}

				mx.Unlock()
			}
			wg.Done()
		}(f, ctx)
	}

	wg.Wait()

	if errCount >= E {
		return ErrErrLimitExceeded
	}

	return nil
}
