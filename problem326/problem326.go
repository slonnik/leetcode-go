package problem326

import (
	"math"
)

func isPowerOfThree(n int) bool {

	one := math.Log10(float64(n))
	two := math.Log10(3)
	result := one / two
	return math.Mod(result, 1.0) < math.Pow10(-10)

}
