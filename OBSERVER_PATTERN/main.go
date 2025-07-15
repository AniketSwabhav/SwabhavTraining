package main

import "fmt"

type Observer interface {
	Update(temp float64, humidity float64, pressure float64)
}

type Subject interface {
	RegisterObserver(o Observer)
	RemoveObserver(o Observer)
	NotifyObservers()
}

type WeatherData struct {
	observers []Observer
	temp      float64
	humidity  float64
	pressure  float64
}

func (w *WeatherData) RegisterObserver(o Observer) {
	w.observers = append(w.observers, o)
}

func (w *WeatherData) RemoveObserver(o Observer) {
	for i, observer := range w.observers {
		if observer == o {
			w.observers = append(w.observers[:i], w.observers[i+1:]...)
			break
		}
	}
}

func (w *WeatherData) NotifyObservers() {
	for _, observer := range w.observers {
		observer.Update(w.temp, w.humidity, w.pressure)
	}
}

func (w *WeatherData) SetMeasurements(temp, humidity, pressure float64) {
	w.temp = temp
	w.humidity = humidity
	w.pressure = pressure
	w.NotifyObservers()
}

type CurrentConditionsDisplay struct {
	temperature float64
	humidity    float64
}

func (c *CurrentConditionsDisplay) Update(temp float64, humidity float64, pressure float64) {
	c.temperature = temp
	c.humidity = humidity
	c.Display()
}

func (c *CurrentConditionsDisplay) Display() {
	fmt.Printf("Current conditions: %.1f°C and %.1f%% humidity\n", c.temperature, c.humidity)
}

func main() {
	weatherData := &WeatherData{}
	currentDisplay := &CurrentConditionsDisplay{}

	weatherData.RegisterObserver(currentDisplay)

	weatherData.SetMeasurements(25.0, 60.0, 1012.0)
	weatherData.SetMeasurements(27.5, 65.0, 1011.0)

	weatherData.RemoveObserver(currentDisplay)

	weatherData.SetMeasurements(30.0, 70.0, 1010.0)
}
