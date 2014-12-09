package xenserver

import (
	"github.com/mitchellh/multistep"
	"time"
)

type InterruptibleWait struct {
	Predicate func () (result bool, err error)
	PredicateInterval time.Duration
	Timeout time.Duration
}

type TimeoutError struct { }
func (err TimeoutError) Error() string {
	return "Timed out"
}

type InterruptedError struct { }
func (err InterruptedError) Error() string {
	return "Interrupted"
}

type PredicateResult struct {
	complete bool
	err error
}

func (wait InterruptibleWait) Wait(state multistep.StateBag) error {
	predicateResult := make(chan PredicateResult, 1)
	stopWaiting := make(chan struct{})
	defer close(stopWaiting)

	go func() {
		for {
			select {
			case <-time.After(wait.PredicateInterval):
				if complete, err := wait.Predicate(); err != nil || complete {
					predicateResult <- PredicateResult{complete, err}
					return
				}

			case <-stopWaiting:
				return
			}
		}
	}()

	timeout := time.After(wait.Timeout)
	for {
		// wait for either install to complete/error,
		// an interrupt to come through, or a timeout to occur
		select {
		case result := <-predicateResult:
			return result.err

		case <-time.After(1 * time.Second):
			if _, ok := state.GetOk(multistep.StateCancelled); ok {
				return InterruptedError{}
			}

		case <-timeout:
			return TimeoutError{}
		}
	}
}
