package square
import (
	"fmt"
	"math"
)
// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)
type myType int
const(
	SidesTriangle=3
	SidesSquare=4
	SidesCircle=0
)
func CalcSquare(sideLen float64, sidesNum myType) float64 {
	if sidesNum==SidesTriangle {
		return Square:=sideLen*sideLen*math.Sqrt(3)/2
		fmt.Println(Square)
	}else if sidesNum==SidesSquare{
		return Square:=sideLen*sideLen
		fmt.Println(Square)
	}else if sidesNum==SidesCircle{
		return Square:=2*math.Pi*sideLen*sideLen
		fmt.Println(Square)
	}else{
		return 0
	}
}

