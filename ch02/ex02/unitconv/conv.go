package unitconv

func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9.0/5.0 + 32) }
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32.0) * 5.0 / 9.0) }

func KToF(k Kelvin) Fahrenheit { return CToF(KToC(k)) }
func FToK(f Fahrenheit) Kelvin { return CToK(FToC(f)) }

func KToC(k Kelvin) Celsius { return Celsius(k) + AbsoluteZeroC }
func CToK(c Celsius) Kelvin { return Kelvin(c - AbsoluteZeroC) }

func FtToM(c Feet) Meter { return Meter(c * 0.3048) }
func MToFt(m Meter) Feet { return Feet(m / 0.3048) }

func LbToKg(p Pound) Kilogram { return Kilogram(p * 0.45359237) }
func KgToLb(k Kilogram) Pound { return Pound(k / 0.45359237) }
