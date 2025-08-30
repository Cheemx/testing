package main

import (
	"fmt"
	"net/http"
	"time"
)

var tenSecondTimeout = 10 * time.Second

func Racer(a, b string) (winner string, error error) {
	return ConfigurableRacer(a, b, tenSecondTimeout)
}

func ConfigurableRacer(a, b string, timeout time.Duration) (winner string, error error) {
	// select basically listens on multiple channels and resolves the case
	// which gets a receive on it first!
	// Since in our ping function we're not actually sending anything on the
	// channel but just closing it when our task is completed which will unblock
	// a receive on our receiver side returning from the function immediately.
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	// time.After() is a very handy function when using select.
	// although not here but there can be a case where code blocks forever
	// because the channels select is listening on never return a value.
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
	}
}

func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
		// Allow me to elaborate on the part
		// where we're not sending an empty struct
		// on the channel and just closing it

		// Why ?
		// Basically a Receive operation on a closed channel
		// unblocks immediately and returns the zero value of the channel's type.

		// Rather if we've sent an empty struct then on the receive
		// side we'd have had to listen on the channel blocking the execution!

		// So closing the channel is a neat trick here and I'm impressed!
	}()
	return ch
}

func measureResponseTime(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	return time.Since(start)
}
