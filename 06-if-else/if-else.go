package main

import f "fmt"


func main () {
    
    // inf else
    if 8%2 == 0 {
        f.Println("8 is even")
    } else {
        f.Println("Number is odd ")
    }


    if i:= 2; i < 1 {
        f.Println(i, "Is less than 1")
    } else {
        f.Println(i, "Is greater than 1")
    }


    // static condition
    if i:= 3; i > 2 {
        f.Println(i, "Is greater than 2")
    } else if i == 3 {
        f.Println(i, "is equal to 3")
    } else {
        f.Println(i, "Is greater than 3")
    }


    // condition change 
    if i:= 3; i > 2 {
        f.Println(i, "Is greater than 2")
    } else if i:= 3; i == 3 {
        f.Println(i, "is equal to 3")
    } else {
        f.Println(i, "Is greater than 3")
    }

}