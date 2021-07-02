package main

import f "fmt"
import m "math"

const s string = "This is constant"

func main () {
    
    const a int = 3
    f.Println(a)
    // need to use alias f  fmt will be undefined`
    f.Println(s)
    
    // number can be represent with _
    const n = 500_000_000
    f.Println(n)

    // scientific num allowed
    const d = 3e20 / n
    f.Println(d)

    // type cast the int`
    f.Println(int64(d))

    // math function
    f.Println(m.Sin(n))

}
