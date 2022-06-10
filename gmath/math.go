// Package gmath includes useful math funcs for games, graphics, and GUIs.
package gmath

import "github.com/seanpfeifer/rigging/num"

// Clamp will return the value clamped between min and max, inclusive.
func Clamp[N num.Real](val, min, max N) N {
	if val < min {
		return min
	} else if val > max {
		return max
	}
	return val
}

// Lerp linearly interpolates between a and b using t. t is clamped between 0.0 and 1.0.
// eg, when t is 0.0, this returns a. When t is 1.0 this returns b. When t is 0.5 this returns the midpoint between a and b.
func Lerp[N num.Real, F num.Float](a, b N, t F) N {
	// If you're using floats, there's effectively no cost to the casting done here with F(b-a).
	// If you're not, you'd have to cast at some point anyway or otherwise do something clever to get a useful interpolation.
	return a + N(F(b-a)*Clamp(t, 0, 1))
}
