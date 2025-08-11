package tummy_test

import (
	"fmt"
	"time"

	"github.com/rakunlabs/tummy"
)

func Example() {
	tummy.Enable()

	tummy.Resume() // it is already resuming

	tummy.SetTime(time.Date(1919, 5, 19, 12, 0, 0, 0, time.UTC))

	tummy.AddDuration(2 * time.Hour)
	fmt.Println(tummy.Now().Hour()) // 14

	tummy.AddDate(0, 0, 1)
	fmt.Println(tummy.Now().Day()) // 20

	tummy.Pause()
	tummy.SetTime(time.Date(1923, 10, 23, 12, 0, 0, 0, time.UTC))

	fmt.Println(tummy.Now().Format(time.RFC3339Nano) == "1923-10-23T12:00:00Z") // true

	tummy.Pause()
	fmt.Println(tummy.IsPaused()) // true
	tummy.Resume()

	tummy.Disable()                                                                  // reset tummy.Now to time.Now
	fmt.Println(tummy.Now().Format(time.RFC3339) == time.Now().Format(time.RFC3339)) // true

	// Output:
	// 14
	// 20
	// true
	// true
	// true
}
