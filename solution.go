package square
import (
	"math"
)

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

type SidesType int

const (
	SidesTriangle = 3
	SidesSquare = 4
	SidesCircle = 0
)

func CalcSquare(sideLen float64, sidesNum SidesType) float64 {
	if sidesNum == SidesTriangle {
		return float64(math.Sqrt(3) / 4 * sideLen * sideLen)
	} else if sidesNum == SidesSquare {
		return float64(sideLen * sideLen)
	} else if sidesNum == SidesCircle {
		return float64(math.Pi * sideLen * sideLen)
	}

	return float64(0)
}
