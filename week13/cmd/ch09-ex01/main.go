package main

import "fmt"

type Kilometers float64
type Meters float64
type Miles float64

func (km Kilometers) ToMiles() Miles {
	return Miles(km / 1.609)
}
func (m Meters) ToMiles() Miles {
	return Miles(m / 1609)
}

func main() {
	kmph := Kilometers(150)
	fmt.Printf("%0.2f Kilometers per hour equals %0.2f Mile per hour\n", kmph, kmph.ToMiles())
	meter := Meters(151000)
	fmt.Printf("%0.2f Meters equals %0.2f Miles\n", meter, meter.ToMiles())
}
