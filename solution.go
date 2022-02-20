package square

import (
	"math"
)

// Define custom int type to hold sides number and update CalcSquare signature
// by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.
// Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

type SidesNumber int

const (
	SidesCircle   = 0
	SidesTriangle = 3
	SidesSquare   = 4
)

func CalcSquare(sideLen float64, sidesNum SidesNumber) float64 {
	switch sidesNum {
	case SidesCircle:
		return math.Pi * sideLen * sideLen
	case SidesTriangle:
		return math.Sqrt(3) / 4 * sideLen * sideLen
	case SidesSquare:
		return sideLen * sideLen
	}
	return 0
}

// func main() {
// 	fmt.Println(CalcSquare(10.0, SidesCircle))
// 	fmt.Println(CalcSquare(10.0, SidesSquare))
// 	fmt.Println(CalcSquare(10.0, SidesTriangle))

// }
