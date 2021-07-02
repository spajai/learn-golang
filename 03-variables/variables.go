package main

import f "fmt";

func main () {

    var a = "Test variable"
    f.Println(a)

    // multiple varibles in single line
    var b,c int = 3,4
    f.Println(b+c)

    // Boolean value`
    var d = true
    f.Println(d)

    // integer assignment
    var e int
    e = 90
    f.Println(e)

    // inline
    g:= "Test Inline assinment and declaration"
    f.Println(g)
 
    // this becomes invalid over write f
    // f:="Apple"
    // f.Println(f)

    // String
    var str string
    str = "asds"
    f.Println(str)

}