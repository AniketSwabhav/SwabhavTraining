package util

import "fmt"

func HandlePanic() {
	if r := recover(); r != nil {
		fmt.Println("recovered from panic:", r)
	}
}
