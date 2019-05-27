// Package convert converts each number into units like temperature in Celsius and Fahrenheit, length in feet and meters, weight in pounds and kilograms
package convert

import "fmt"

type Celsius float64
type Fahrenheit float64
type Meter float64
type Feet float64
type Pound float64
type Kilogram float64

const (
	AbsoluteZeroC      Celsius  = -273.15
	OneMeterInFeet     Feet     = 3.28084
	OneFeetInMeter     Meter    = 1.000000032
	OnePoundInKilogram Kilogram = 0.453592
	OneKilogramInPound Pound    = 0.99999918429
)

func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }
func (m Meter) String() string      { return fmt.Sprintf("%g meter", m) }
func (f Feet) String() string       { return fmt.Sprintf("%g feet", f) }
func (p Pound) String() string      { return fmt.Sprintf("%g pound/s", p) }
func (k Kilogram) String() string   { return fmt.Sprintf("%g kilogram/s", k) }

// CToF converts a Celsius temperature to Fahrenheit.
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// FToC converts a Fahrenheit temperature to Celsius.
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

// MToF converts a Meter length to Feet.
func MToF(m Meter) Feet { return Feet(Feet(m) * OneMeterInFeet) }

// FToM converts a Feet length to Meter.
func FToM(f Feet) Meter { return Meter(f / OneMeterInFeet) }

// PToK converts a Pund weight to Kilogram.
func PToK(p Pound) Kilogram { return Kilogram(float64(p) * float64(OnePoundInKilogram)) }

// KToP converts a Pund weight to Kilogram.
func KToP(k Kilogram) Pound { return Pound(k / OnePoundInKilogram) }
