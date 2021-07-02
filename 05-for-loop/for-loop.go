package main

import f "fmt"


func main () {

    i:= 1

    // simple loop
    for i < 10 {
        f.Println("Printing i: ",i)
        i++
    }

    // classic for
    for j := 1; j < 100; j++ {
        f.Println("printing j: ", j)
    }

    // loop break
    for {
        f.Println("loop")
        break
    }


   for n := 0; n <= 5; n++ {
        if n%2 == 0 {
            continue
        }
        f.Println(n)
   }


}