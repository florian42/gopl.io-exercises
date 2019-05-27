package convert

import (
	"math"
	"testing"
)

func TestTempConv(t *testing.T) {
	tests := []struct {
		f Fahrenheit
		c Celsius
	}{
		{68, 20},
		{32, 0},
		{-40, -40},
	}
	eps := 0.0000001 // acceptable floating point error
	for _, test := range tests {
		if math.Abs(float64(CToF(test.c)-test.f)) > eps {
			t.Errorf("CToF(%s): got %s, want %s", test.c, CToF(test.c), test.f)
		}
		if math.Abs(float64(FToC(test.f)-test.c)) > eps {
			t.Errorf("FToC(%s): got %s, want %s", test.f, FToC(test.f), test.c)
		}
	}
}

func TestLengthConv(t *testing.T) {
	tests := []struct {
		m Meter
		f Feet
	}{
		{68, 223.097},
		{32, 104.987},
		{250, 820.21},
	}
	eps := 0.01 // acceptable floating point error
	for _, test := range tests {
		if math.Abs(float64(FToM(test.f)-test.m)) > eps {
			t.Errorf("FToM(%s): got %s, want %s", test.f, FToM(test.f), test.m)
		}
		if math.Abs(float64(MToF(test.m)-test.f)) > eps {
			t.Errorf("MToF(%s): got %s, want %s", test.m, MToF(test.m), test.f)
		}
	}
}

func TestWeightConv(t *testing.T) {
	tests := []struct {
		k Kilogram
		p Pound
	}{
		{68, 149.914},
		{32, 70.5479},
		{250, 551.156},
	}
	eps := 0.01 // acceptable floating point error
	for _, test := range tests {
		if math.Abs(float64(KToP(test.k)-test.p)) > eps {
			t.Errorf("KToP(%s): got %s, want %s", test.k, KToP(test.k), test.p)
		}
		if math.Abs(float64(PToK(test.p)-test.k)) > eps {
			t.Errorf("PToK(%s): got %s, want %s", test.p, PToK(test.p), test.k)
		}
	}
}
