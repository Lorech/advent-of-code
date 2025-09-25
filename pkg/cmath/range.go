package cmath

import "math"

// Finds the closest number in [targetLower, targetUpper] given an input
// number within [domainLower, domainUpper]. Distances are computed circularly
// over the domain range.
func ClosestInRange(num, domainLower, domainUpper, targetLower, targetUpper int) int {
	// Normalize inputs to be inside the domain
	if num < domainLower {
		num = domainLower
	}
	if num > domainUpper {
		num = domainUpper
	}

	domainSize := domainUpper - domainLower + 1
	best := targetLower
	minDist := math.MaxInt

	for i := targetLower; i <= targetUpper; i++ {
		// Direct distance
		dist := int(math.Abs(float64(num - i)))

		// Wrap-around distance
		wrapDist := domainSize - dist

		// Pick smaller of the two
		effectiveDist := min(wrapDist, dist)
		if effectiveDist < minDist {
			minDist = effectiveDist
			best = i
		}
	}

	return best
}
