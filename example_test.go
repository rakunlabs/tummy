package tummy_test

import (
	"fmt"
	"time"

	"github.com/rakunlabs/tummy"
)

func Example() {
	tummy.Enable()

	tummy.SetTime(time.Date(1919, 5, 19, 12, 0, 0, 0, time.UTC))

	tummy.AddDuration(2 * time.Hour)
	fmt.Println(tummy.Now().Hour()) // 14

	tummy.AddDate(0, 0, 1)
	fmt.Println(tummy.Now().Day()) // 20

	// Output:
	// 14
	// 20
}
