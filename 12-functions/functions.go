package main

// a,b with int and return int
func add (a int,b int) int {
    return a + b
}

// abc with type int and return type is int
func add1 (a,b,c int) int {
    // lets call other func
    return add(a,b) + c
}

// end func
func end () {
    print ("\nProgram has beeen ended calling end\n")
}

func main () {
    res := add(1,4)

    print("\n1+4 =", res)
    res1 := add1(4,5,1)
    print("\n4 + 5 + 1=", res1)
    end()
}