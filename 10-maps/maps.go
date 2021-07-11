package main

import "fmt"

func main () {

    // this will create a map of String keys and values as integer
    // make(map[key-type]val-type)
    m := make(map[string]int)

    m["k1"] = 78
    m["j1"] = 56

    fmt.Println("Map:", m)
    // output
    // Map: map[j1:56 k1:78]

    // this will get the value
    v1 := m["k1"]
    fmt.Println("v1: ", v1)

    // length
    fmt.Println("len:", len(m))

    // delete the k2 key
    // deleting invalid key will not error/warn
    delete(m, "k2")
    fmt.Println("map:", m)


    // The optional second return value when getting a value from a map indicates if the key was present in the map. This can be used to disambiguate between missing keys and keys with zero values like 0 or "". Here we didn’t need the value itself, so we ignored it with the blank identifier _
    _, prs := m["k2555"]
    // return true/false
    fmt.Println("prs:", prs)


    // oneliner
    n := map[string]int{"1": 1, "bar": 2}
    fmt.Println("map:", n)


    m1 := make(map[int]int)
    for i := 0; i < 10; i++ {
        m1[i] = i+ 10
    }
    fmt.Println("map:", m1)


}