// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 41.

//!+

package tempconv

// CToF converts a Celsius temperature to Fahrenheit.
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// KToF converts a Kelvin temperature to Fahrenheit
func KToF(k Kelvin) Kelvin { return Kelvin(k*9/5 + 32) }

// KToC converts a Kelvin temperature to Celsius
func KToC(k Kelvin) Celsius { return Celsius(k + Kelvin(AbsoluteZerok)) }

// FToC converts a Fahrenheit temperature to Celsius.
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

// FToK converts a Fahrenheit temperature to Celsius.
func FToK(f Fahrenheit) Kelvin { return Kelvin((f - 32) * 5 / 9) }

//!-
