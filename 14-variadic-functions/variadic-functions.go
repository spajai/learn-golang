package main

import "fmt"

// this function takes any number of input
// if return is specified then the function defination expect the return type strictly now function dont return anything
func sum_any (nums...int) {
	fmt.Println(nums," ")
	total := 0

	for _,i := range nums {
		total += i
	}
	fmt.Println("Total",total)
}

func main () {
	// send list
	sum_any(1,3,4,5)

	// invalid operation because the function dont return anything
	// b := sum_any(34,23,23,23,54)
	// print("b",b)
	nums := []int{1,2,3,45,56,67}
	// send array
	sum_any(nums...)
}