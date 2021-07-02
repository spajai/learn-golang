package main

import "fmt"

func main () {

    // this will add in new line
    fmt.Println("go" + "lang" + "String" + "concat")
    // this will not add new line
    print("go" + "lang" + "String" + "concat")

    fmt.Println("\n1+1=",1+1)
    // manual new line
    print("1+1=", 1+1, "\n")

    fmt.Println("7/3.0 = ", 7.0/3.0)
    print("7/3.0 = ", 7.0/3.0, "\n")

    // will print true or false
    fmt.Println(true && false)
    fmt.Println(true || false)
    fmt.Println(!true)

}