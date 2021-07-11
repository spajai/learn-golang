package main
import "fmt"

// This `fact` function calls itself until it reaches the
// base case of `fact(0)`.
func fact (n int64) int64 {
	if n == 0 {
		return 1
	}
	return (int64(n) * fact(int64(n-1)))
}
func main () { 
	fmt.Println(fact(30))
}