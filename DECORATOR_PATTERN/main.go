package main

import "fmt"

//Interfcae
type Beverage interface {
	Description() string
	Cost() float32
}

// ------------------------------------------------------
// concrete component
type Espresso struct{}

func (e *Espresso) Description() string {
	return "Espresso"
}

func (e *Espresso) Cost() float32 {
	return 1.99
}

// ------------------------------------------------------
// concrete component
type HouseBlend struct{}

func (h *HouseBlend) Description() string {
	return "House Blend Coffee"
}

func (h *HouseBlend) Cost() float32 {
	return 0.89
}

// ------------------------------------------------------
//  decorator
type CondimentDecorator struct {
	Beverage Beverage
}

// ------------------------------------------------------
// Mocha is a concrete decorator
type Mocha struct {
	Beverage Beverage
}

func (m *Mocha) Description() string {
	return m.Beverage.Description() + ", Mocha"
}

func (m *Mocha) Cost() float32 {
	return m.Beverage.Cost() + 0.20
}

// Whip is concrete decorator
type Whip struct {
	Beverage Beverage
}

func (w *Whip) Description() string {
	return w.Beverage.Description() + ", Whip"
}

func (w *Whip) Cost() float32 {
	return w.Beverage.Cost() + 0.10
}

func Bill(myDrink Beverage) {
	fmt.Println("Order:", myDrink.Description())
	fmt.Printf("Total: $%.2f\n", myDrink.Cost())
	fmt.Println("-------------------------")
}

func main() {
	var myDrink Beverage

	myDrink = &Espresso{}
	Bill(myDrink)

	myDrink = &Mocha{Beverage: myDrink}
	Bill(myDrink)

	myDrink = &Whip{Beverage: myDrink}
	Bill(myDrink)

}
