package schedule

import (
	"fmt"
	"time"
)

func ExampleRepeat() {
	doSomething := func() { fmt.Print("*") }
	stop := Repeat(doSomething, time.Second)
	time.Sleep(3 * time.Second)
	stop <- true
	// Output:
	// ***
}
