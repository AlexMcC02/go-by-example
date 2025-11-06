package main

import (
	"fmt"
	"time"
)

// Go's built-in timer feature allows the execution of code at some point in the future,
// or repeatedly at some internval.

func main() {

	// Timers represent a single point in the future. You specify both the amount of time
	// and the channel that will be notified when the specified time elapses.
	timer1 := time.NewTimer(2 * time.Second)

	<-timer1.C
	fmt.Println("Timer 1 fired")

	timer2 := time.NewTimer(time.Second)
	
	// If you just need to wait, time.Sleep can be used. A timer may be useful in that
	// you can cancel the timer before it fires.
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 fired")
	}()

	stop2 := timer2.Stop()

	if stop2 {
		fmt.Println("Timer 2 stopped")
	}

	time.Sleep(2 * time.Second)
}