package main

import (
	"fmt"
	"sync"
)

// Singleton: ChocolateBoiler
type ChocolateBoiler struct {
	empty  bool
	boiled bool
}

var instance *ChocolateBoiler
var once sync.Once

// GetInstance ensures only one boiler ever exists
func GetInstance() *ChocolateBoiler {
	once.Do(func() {
		fmt.Println("Creating the single instance of ChocolateBoiler...")
		instance = &ChocolateBoiler{empty: true, boiled: false}
	})
	return instance
}

// Fill the boiler
func (b *ChocolateBoiler) Fill() {
	if b.empty {
		fmt.Println("Filling the boiler with milk and chocolate...")
		b.empty = false
		b.boiled = false
	} else {
		fmt.Println("Boiler is already filled.")
	}
}

// Boil the contents
func (b *ChocolateBoiler) Boil() {
	if !b.empty && !b.boiled {
		fmt.Println("Boiling the contents...")
		b.boiled = true
	} else {
		fmt.Println("Cannot boil: either already boiled or empty.")
	}
}

// Drain the boiled mixture
func (b *ChocolateBoiler) Drain() {
	if !b.empty && b.boiled {
		fmt.Println("Draining the boiled chocolate...")
		b.empty = true
	} else {
		fmt.Println("Cannot drain: either not boiled or already empty.")
	}
}

func main() {
	// Accessing singleton
	boiler := GetInstance()

	boiler.Fill()
	boiler.Boil()
	boiler.Drain()

	// Try accessing again
	boiler2 := GetInstance()
	if boiler == boiler2 {
		fmt.Println("Same instance reused.")
	}
}
