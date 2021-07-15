package main
import "fmt"

// define the struct
type rect struct {
	height, width int
}

// write the function area that takes the input as pointer of varible struct type of rect 
// and return the int 
//function will be regular
func area(d *rect) int {
	// access the  property height and width
	return d.height * d.width
}

// 
func (r *rect) area1 () int {
	return r.height * r.width
}
func (r *rect) peri () int {
	return 2*(r.height + r.width)
}

func main () { 
	m := rect{height: 10  , width : 9}
	// pass the pointer of type rect
	//regular function call
	fmt.Println("area: ", area(&m))

	//function call with note the syntax is different
	fmt.Println("area: ", m.area1())

	//assignment of pointert to m into rp
	rp := &m
	fmt.Println("area: ", rp.peri())
}
