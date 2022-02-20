package square

import "math"

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#
type A int

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

const (
	Square   A = 4
	Triangle A = 3
	Circle   A = 0
)

func CalcSquare(sideLen float64, sidesNum A) float64 {
	switch sidesNum {
	case Square:
		return math.Sqrt(sideLen * sideLen)
	case Triangle:
		return math.Sqrt(math.Sqrt(3) / 4 * (sideLen * sideLen))
	case Circle:
		return math.Sqrt(math.Pi * (sideLen * sideLen))
	default:
		return 0
	}
}
