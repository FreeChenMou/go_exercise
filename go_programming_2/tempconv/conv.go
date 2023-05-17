package tempconv

// CToF converts a Celsius temperature to Fahrenheit.
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// FToC converts a Fahrenheit temperature to Celsius.
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

// MToI converts a Meter temperature to Inch.
func MToI(m Meter) Inch {
	return Inch(m * 39.3701)
}

// IToM converts a Inch temperature to Meter.
func IToM(inch Inch) Meter {
	return Meter(inch / 39.3701)
}

// KToP converts a Kilogram temperature to Pound.
func KToP(kg Kilogram) Pound {
	return Pound(kg * 2.2046)
}

// PToK converts a Pound temperature to Kilogram.
func PToK(p Pound) Kilogram {
	return Kilogram(p / 2.2046)
}
