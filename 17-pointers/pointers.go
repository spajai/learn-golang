package main
import "fmt"

func zeroval (i_val int) {
	i_val = 0
}

func zeroptr (iptr *int) {
	*iptr = 0
}

func main () { 
	i:= 1
	fmt.Println("Initial: ", i)

	zeroval(i)
	fmt.Println("zeroval: ", i)
	
	zeroptr(&i)
	fmt.Println("zeroptr: ",i)
	
	i= 15
	zeroptr(&i)
	fmt.Println("zeroptr: ",i)
	
}