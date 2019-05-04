package schedule

import "time"

func Repeat(action func(), delay time.Duration) chan bool {
	stop := make(chan bool)
	go func() {
		for {
			select {
			case <-time.After(delay):
			case <-stop:
				return
			}
			action()
		}
	}()
	return stop
}
