package main

import "fmt"

func multi() (int,string) {
    return 7, "abc"
}

func main () {
    a,b := multi()
    fmt.Println(a,b)

    _,v := multi()
    fmt.Println(v)

    v1,_ := multi()
    fmt.Println(v1)
}