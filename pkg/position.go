package picture

//Pos Position X AND Y
type Pos struct {
	X, Y float32
}

//NewPos Return a Pos truct
func NewPos(x, y float32) Pos {
	return Pos{x, y}
}

//Lerp Find distance between two point
func Lerp(a float32, b float32, pct float32) float32 {
	return a + pct*(b-a)
}
