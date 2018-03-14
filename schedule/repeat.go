package schedule

import "time"

//function: Non-blocking loop timer, do action per delay time
//return: stop control
func Repeat(action func(), delay time.Duration) chan bool {
	stop := make(chan bool)
	go func() {
		for {
			action()
			select {
			case <-time.After(delay):
			case <-stop:
				return
			}
		}
	}()
	return stop
}
