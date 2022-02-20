package square

import (
	"math"
)

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#
type A int

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

const (
	SidesSquare   A = 4
	SidesTriangle A = 3
	SidesCircle   A = 0
)

func CalcSquare(sideLen float64, sidesNum A) float64 {
	switch sidesNum {
	case SidesSquare:
		return math.Pow(sideLen, 2)
	case SidesTriangle:
		return (sideLen * sideLen) * (math.Sqrt(3) / 4)
	case SidesCircle:
		return math.Pi * (sideLen * sideLen)
	default:
		return 0
	}
}
