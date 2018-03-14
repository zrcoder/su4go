package schedule

import (
	"fmt"
	"time"
	"testing"
)

func ExampleRepeat() {
	doSomething := func() { fmt.Print("*") }
	stop := Repeat(doSomething, time.Second)
	time.Sleep(3 * time.Second)
	stop <- true
}

func TestRepeat(t *testing.T) {
	ExampleRepeat()
}
