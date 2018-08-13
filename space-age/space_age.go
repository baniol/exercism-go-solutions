// Package space implements Age function for calculating age
// on planets of the Solar System
package space

type Planet string

const earthPeriod float64 = 31557600

var planets = map[Planet]float64{
	"Earth":   1,
	"Mercury": 0.2408467,
	"Venus":   0.61519726,
	"Mars":    1.8808158,
	"Jupiter": 11.862615,
	"Saturn":  29.447498,
	"Uranus":  84.016846,
	"Neptune": 164.79132,
}

// Age recalculates the input age on a given planet to Earth age in years
func Age(seconds float64, planet Planet) float64 {
	if _, ok := planets[planet]; !ok {
		panic("Planet not found in the Solar System")
	}
	return seconds / earthPeriod / planets[planet]
}
