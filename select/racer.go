package racer

import (
	"fmt"
	"net/http"
	"time"
)

const tenSecondTimeout = 10 * time.Second

func Racer(urlOne string, urlTwo string) (winner string, err error) {
	return ConfigurableRacer(urlOne, urlTwo, tenSecondTimeout)
}

func ConfigurableRacer(urlOne string, urlTwo string, timeout time.Duration) (winner string, err error) {
	select {
	case <-ping(urlOne):
		return urlOne, nil
	case <-ping(urlTwo):
		return urlTwo, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waiting for %s and %s", urlOne, urlTwo)
	}
}

func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}
