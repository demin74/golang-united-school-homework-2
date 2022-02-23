package square

import "math"

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#
type CountSides int

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
const (
 		SidesCircle		CountSides 	= iota    //0
 		SidesTriangle 				= iota+2  //3=1+2
		SidesSquare					= iota+2  //4=2+2
	  )

var Q int64=0
var Qf float64=0
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

func CalcSquare(sideLen float64, sidesNum CountSides) float64 {
///////////////////////////////////////////////////////////////////////////////
	Q:=sideLen*sideLen // квадрат числа (возведение в вторую степень - число степени два и т.д. и т.п.)
	 Qf:=float64(Q)

	if sidesNum==4 {
	 return Qf
	} else if sidesNum==3 {
	 return Qf*(math.Sqrt(3)/4)
	} else if sidesNum==0 {
	 return Qf*math.Pi
	} else {
	 return 0
	}
///////////////////////////////////////////////////////////////////////////////	
}
