package main

import "fmt"

type Celcius float64
type Fahrenheit float64

const (
	AbsoluteZeroC Celcius = -273.15
	Freezing      Celcius = 0
	Boiling       Celcius = 100
)

func CToF(c Celcius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }
func FtoC(f Fahrenheit) Celcius { return Celcius((f - 32) * 5 / 9) }

// Association with the Celcius type
func (c Celcius) String() string { return fmt.Sprintf("%g°C", c) }

func main() {
	c := FtoC(212.0)
	fmt.Println(c.String())
}
