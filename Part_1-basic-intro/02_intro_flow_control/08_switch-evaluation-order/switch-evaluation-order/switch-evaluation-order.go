package main

import (
	"fmt"
	"time"
)

func main() {
	today := time.Now().Weekday()

	// Uncomment to adjust the date and test
	// today := time.Date(2023, 02, 9, 10, 10, 0, 0, time.UTC).Weekday()


	switch time.Saturday {
	case today:
		fmt.Println("Today is Saturday!")
	case today + 1:
		fmt.Println("Tomorrow is Saturday!")
	case today + 2:
		fmt.Println("In two days time it'll be Saturday")
	default:
		fmt.Println("Saturday is still far away")
	}

	// Alternatively
	// switch today {
	// case time.Saturday:
	// 	fmt.Println("Today is Saturday!")
	// case time.Friday:
	// 	fmt.Println("Tomorrow is Saturday!")
	// case time.Thursday:
	// 	fmt.Println("In two days time it'll be Saturday")
	// default:
	// 	fmt.Println("Saturday is still far away")
	// }
}