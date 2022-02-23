package square

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

type sideCountType int

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

const (
	SidesCircle   sideCountType = 0
	SidesTriangle sideCountType = 3
	SidesSquare   sideCountType = 4
)

func CalcSquare(sideLen float64, sidesNum sideCountType) float64 {

	switch sidesNum {
	case SidesSquare:
		res = math.Pow(sideLen, 2.0)
	case SidesTriangle:
		res = (math.Pow(sideLen, 2.0) * math.Sqrt(3.0)) / 4
	case SidesCircle:
		res = math.Pi * math.Pow(sideLen, 2.0)
	default:
		res = 0
	}
	return res
}
