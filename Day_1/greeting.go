package main

import (
	"fmt"
	"time"
)

func main() {

	var currentTime = time.Now()

	if currentTime.Hour() > 6 && currentTime.Hour() < 11 {
		fmt.Println("Good Morning!")
	} else if currentTime.Hour() > 11 && currentTime.Hour() < 16 {
		fmt.Println("Good Afternoon!")
	} else if currentTime.Hour() > 16 && currentTime.Hour() < 21 {
		fmt.Println("Good Evening!")
	} else if currentTime.Hour() >= 21 || currentTime.Hour() < 6 {
		fmt.Println("Good Night!")
	} else {
		fmt.Println("Not validd time!")
	}

}
