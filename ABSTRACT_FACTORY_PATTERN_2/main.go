package main

// IngredientFactory interface         // Abstract Factory
//  ├── NYIngredientFactory struct     // Concrete Factory
//  └── ChicagoIngredientFactory struct

// Ingredients:
//  ├── Dough
//  ├── Sauce
//  └── Cheese

// Pizza struct (depends on IngredientFactory)
//  └── CheesePizza (uses ingredients via factory)

import "fmt"

//
// INGREDIENT INTERFACES
//
type Dough interface {
	GetDough() string
}
type Sauce interface {
	GetSauce() string
}
type Cheese interface {
	GetCheese() string
}

//
// INGREDIENT FACTORY INTERFACE
//
type IngredientFactory interface {
	CreateDough() Dough
	CreateSauce() Sauce
	CreateCheese() Cheese
}

//
// NY INGREDIENTS
//
type NYDough struct{}

func (d *NYDough) GetDough() string { return "Thin Crust Dough" }

type NYSauce struct{}

func (s *NYSauce) GetSauce() string { return "Marinara Sauce" }

type NYCheese struct{}

func (c *NYCheese) GetCheese() string { return "Reggiano Cheese" }

type NYIngredientFactory struct{}

func (f *NYIngredientFactory) CreateDough() Dough   { return &NYDough{} }
func (f *NYIngredientFactory) CreateSauce() Sauce   { return &NYSauce{} }
func (f *NYIngredientFactory) CreateCheese() Cheese { return &NYCheese{} }

//
// CHICAGO INGREDIENTS
//
type ChicagoDough struct{}

func (d *ChicagoDough) GetDough() string { return "Extra Thick Crust Dough" }

type ChicagoSauce struct{}

func (s *ChicagoSauce) GetSauce() string { return "Plum Tomato Sauce" }

type ChicagoCheese struct{}

func (c *ChicagoCheese) GetCheese() string { return "Shredded Mozzarella" }

type ChicagoIngredientFactory struct{}

func (f *ChicagoIngredientFactory) CreateDough() Dough   { return &ChicagoDough{} }
func (f *ChicagoIngredientFactory) CreateSauce() Sauce   { return &ChicagoSauce{} }
func (f *ChicagoIngredientFactory) CreateCheese() Cheese { return &ChicagoCheese{} }

//
// PIZZA BASE + CHEESE PIZZA
//
type Pizza interface {
	Prepare()
	Bake()
	Cut()
	Box()
	GetName() string
}

type CheesePizza struct {
	name    string
	dough   Dough
	sauce   Sauce
	cheese  Cheese
	factory IngredientFactory
}

func (p *CheesePizza) Prepare() {
	fmt.Println("Preparing", p.name)
	p.dough = p.factory.CreateDough()
	p.sauce = p.factory.CreateSauce()
	p.cheese = p.factory.CreateCheese()
	fmt.Println("Using:", p.dough.GetDough())
	fmt.Println("Using:", p.sauce.GetSauce())
	fmt.Println("Using:", p.cheese.GetCheese())
}
func (p *CheesePizza) Bake()           { fmt.Println("Baking", p.name) }
func (p *CheesePizza) Cut()            { fmt.Println("Cutting", p.name) }
func (p *CheesePizza) Box()            { fmt.Println("Boxing", p.name) }
func (p *CheesePizza) GetName() string { return p.name }

//
// UPDATED PIZZA STORES USING FACTORIES
//
type PizzaStore interface {
	OrderPizza(pizzaType string) Pizza
}

type NYPizzaStore struct{}

func (s *NYPizzaStore) OrderPizza(pizzaType string) Pizza {
	var pizza Pizza
	factory := &NYIngredientFactory{}

	if pizzaType == "cheese" {
		pizza = &CheesePizza{name: "NY Style Cheese Pizza", factory: factory}
	}
	pizza.Prepare()
	pizza.Bake()
	pizza.Cut()
	pizza.Box()
	return pizza
}

type ChicagoPizzaStore struct{}

func (s *ChicagoPizzaStore) OrderPizza(pizzaType string) Pizza {
	var pizza Pizza
	factory := &ChicagoIngredientFactory{}

	if pizzaType == "cheese" {
		pizza = &CheesePizza{name: "Chicago Style Cheese Pizza", factory: factory}
	}
	pizza.Prepare()
	pizza.Bake()
	pizza.Cut()
	pizza.Box()
	return pizza
}

//
// MAIN
//
func main() {
	var store PizzaStore

	fmt.Println("---- NY Order ----")
	store = &NYPizzaStore{}
	pizza := store.OrderPizza("cheese")
	fmt.Println("Ordered:", pizza.GetName())

	fmt.Println("\n---- Chicago Order ----")
	store = &ChicagoPizzaStore{}
	pizza = store.OrderPizza("cheese")
	fmt.Println("Ordered:", pizza.GetName())
}
