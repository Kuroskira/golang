package conv

import "fmt"

type Celsius float64
type Fahrenheit float64
type Kelvin float64
type Meter float64
type Feet float64

const (
	AbsoluteZeroC Celsius = -273.15
	Freezing      Celsius = 0
	Boiling       Celsius = 100
)

// Association with the Celcius type
func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }
func (k Kelvin) String() string     { return fmt.Sprintf("%g°K", k) }
func (m Meter) String() string      { return fmt.Sprintf("%gm", m) }
func (ft Feet) String() string      { return fmt.Sprintf("%gft", ft) }
