package main

import "fmt"

func main () {
    // declare array of length 5
    var a[6] int

    // print the array variable a [0 0 0 0 0]
    fmt.Println(a)

    // print the dynamic array 
    var b[] string
    fmt.Println(b)

    // print length
    fmt.Println(len(a))

    a[4] = 100
    fmt.Println("Set value of 4th index a", a)
    fmt.Println(a[4])

    // following will not work since array defined with 0 lenght
    // b[0] = "100"
    // fmt.Println("Set value of 4th index b ", a)
    // fmt.Println(b[4])
 
    // oneliner array defination
    c := [5]int{1, 2, 3, 4, 5}
    fmt.Println(c)


    // two dimentonal array
    var TwoDimArr[5][5] int

    for i := 0;  i <5 ; i++ {
        for j :=0; j < 5 ; j++ {
            TwoDimArr[i][j] = i + j
        }
    }

    fmt.Println(TwoDimArr);

    x := [5]int{10, 20, 30, 40, 50}   // Intialized with values
    var y [5]int = [5]int{10, 20, 30} // Partial assignment

    fmt.Println(x)
    fmt.Println(y)
}

/*
    Notes:-
        Array have fixed size and 2 dimentional array supported

*/