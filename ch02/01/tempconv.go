package tempconv

import "fmt"

// Celsius type
type Celsius float64

// Fahrenheit type
type Fahrenheit float64

// Kelvins type
type Kelvins float64

const (
	// AbsoluteZeroC in Celsius
	AbsoluteZeroC Celsius = -273.15
	// FreezingC in Celsius
	FreezingC Celsius = 0
	// BoilingC in Celsius
	BoilingC Celsius = 100
	// AbsoluteZeroK in Kelvins
	AbsoluteZeroK Kelvins = 0
)

func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }
func (k Kelvins) String() string {
	if k < 0 {
		return "Not a Kelvins temperature."
	}
	return fmt.Sprintf("%gK", k)
}
