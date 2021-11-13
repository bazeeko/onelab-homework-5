package onelabhomework5

import (
	"errors"
	"sync"
)

var ErrErrLimitExceeded = errors.New("error limit exceeded")

func Execute1(tasks []func() error, E int) error {
	wg := sync.WaitGroup{}
	mx := sync.Mutex{}

	errCount := 0

	for _, f := range tasks {
		wg.Add(1)
		go func(f func() error) {
			err := f()
			if err != nil {
				mx.Lock()
				errCount++
				mx.Unlock()
			}
			wg.Done()
		}(f)
	}

	wg.Wait()

	if errCount < E {
		return nil
	}

	return ErrErrLimitExceeded
}
