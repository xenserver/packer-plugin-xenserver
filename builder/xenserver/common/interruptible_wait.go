package common

import (
	"github.com/mitchellh/multistep"
	"time"
)

type InterruptibleWait struct {
	Timeout time.Duration

	// optional:
	Predicate         func() (result bool, err error)
	PredicateInterval time.Duration
}

type TimeoutError struct{}

func (err TimeoutError) Error() string {
	return "Timed out"
}

type InterruptedError struct{}

func (err InterruptedError) Error() string {
	return "Interrupted"
}

type PredicateResult struct {
	complete bool
	err      error
}

/* Wait waits for up to Timeout duration, checking an optional Predicate every PredicateInterval duration.
   The first run of Predicate is immediately after Wait is called.
   If the command is interrupted by the user, then an InterruptedError is returned.
   If Predicate is not nil, a timeout leads to TimeoutError being returned, and a successful Predicate run leads to nil being returned.
   If Predicate is nil, a timeout is not an error, and nil is returned.
*/
func (wait InterruptibleWait) Wait(state multistep.StateBag) error {
	predicateResult := make(chan PredicateResult, 1)
	stopWaiting := make(chan struct{})
	defer close(stopWaiting)

	if wait.Predicate != nil {
		go func() {
			for {
				if complete, err := wait.Predicate(); err != nil || complete {
					predicateResult <- PredicateResult{complete, err}
					return
				}

				select {
				case <-time.After(wait.PredicateInterval):
					continue
				case <-stopWaiting:
					return
				}
			}
		}()
	}

	timeout := time.After(wait.Timeout)
	for {
		// wait for either install to complete/error,
		// an interrupt to come through, or a timeout to occur

		if _, ok := state.GetOk(multistep.StateCancelled); ok {
			return InterruptedError{}
		}

		select {
		case result := <-predicateResult:
			return result.err

		case <-time.After(1 * time.Second):
			continue

		case <-timeout:
			if wait.Predicate != nil {
				return TimeoutError{}
			} else {
				return nil
			}
		}
	}
}
