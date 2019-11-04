package main

import (
	"fmt"
	"time"
)

func main() {
	i := 2
	fmt.Print("Write ", i, " as ")

	switch i {
	case 1:
		fmt.Println("one\n")
	case 2:
		fmt.Println("two\n")
	case 3:
		fmt.Println("three\n")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend!\n")
	case time.Friday:
		fmt.Println("Welcome to Friday! It's almost the weekend!\n")
	case time.Monday:
		fmt.Println("Someone has a case of the monday's... ;)\n")
	default:
		fmt.Println("It's the middle of the week bud... You'll survive!\n")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's only the beginning of your day.\n")
	case 12 < t.Hour() && t.Hour() < 17:
		fmt.Println("Half way through the day!\n")
	default:
		fmt.Println("Good night!\n")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a boolean.")
		case int:
			fmt.Println("I'm an integer.")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}

	whatAmI(true)
	whatAmI(1)
	whatAmI("heylo")
}
