package main

// Pizza           - interface
// CheesePizza     - struct
// VeggiePizza     - struct

// PizzaStore      - interface
// NYPizzaStore    - struct
// ChicagoPizzaStore - struct

import "fmt"

//
// PIZZA INTERFACE & CONCRETE PIZZAS
//
type Pizza interface {
	Prepare()
	Bake()
	Cut()
	Box()
	GetName() string
}

type CheesePizza struct {
	name string
}

func (p *CheesePizza) Prepare() {
	fmt.Println("Preparing", p.name)
}
func (p *CheesePizza) Bake() {
	fmt.Println("Baking", p.name)
}
func (p *CheesePizza) Cut() {
	fmt.Println("Cutting", p.name)
}
func (p *CheesePizza) Box() {
	fmt.Println("Boxing", p.name)
}
func (p *CheesePizza) GetName() string {
	return p.name
}

//
// PIZZA STORE - FACTORY METHOD PATTERN
//
type PizzaStore interface {
	OrderPizza(pizzaType string) Pizza
	CreatePizza(pizzaType string) Pizza // Factory Method
}

type NYPizzaStore struct{}

func (s *NYPizzaStore) OrderPizza(pizzaType string) Pizza {
	pizza := s.CreatePizza(pizzaType)
	pizza.Prepare()
	pizza.Bake()
	pizza.Cut()
	pizza.Box()
	return pizza
}

func (s *NYPizzaStore) CreatePizza(pizzaType string) Pizza {
	if pizzaType == "cheese" {
		return &CheesePizza{name: "NY Style Cheese Pizza"}
	}
	return nil
}

type ChicagoPizzaStore struct{}

func (s *ChicagoPizzaStore) OrderPizza(pizzaType string) Pizza {
	pizza := s.CreatePizza(pizzaType)
	pizza.Prepare()
	pizza.Bake()
	pizza.Cut()
	pizza.Box()
	return pizza
}

func (s *ChicagoPizzaStore) CreatePizza(pizzaType string) Pizza {
	if pizzaType == "cheese" {
		return &CheesePizza{name: "Chicago Style Cheese Pizza"}
	}
	return nil
}

//
// MAIN
//
func main() {
	var store PizzaStore

	store = &NYPizzaStore{}
	pizza := store.OrderPizza("cheese")
	fmt.Println("Ordered a", pizza.GetName())

	store = &ChicagoPizzaStore{}
	pizza = store.OrderPizza("cheese")
	fmt.Println("Ordered a", pizza.GetName())
}
