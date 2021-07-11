package main

import "fmt"

func main () {
    // var a int = 2
    // b := 9
    // g := make([]int,9)
    // fmt.Println(a,b,g)

    nums := []int{1, 2, 3, 4}

    sum := 0

    // range on arrays and slices provides both the index and value for each entry
    // Above we didn’t need the index, so we ignored it with the blank identifier _. Sometimes we actually want the indexes though.
    for _ , num := range nums {
        sum += num
    }

    fmt.Println(sum)


    for i,num := range nums {
        // this will print the index and value
        fmt.Println(i,num)
        
        if i == 3 {
            fmt.Println("index 3 mached:", i)
        }
    }


    // print key value from map
    kvs := map[string]string{"a": "apple", "b": "banana"}

    for k,v := range kvs {
            fmt.Printf("Key %s, values %s  direct values %s\n", k,v,kvs[k])
    }

    // print keys from map
    for k := range kvs {
        fmt.Println("key:", k)
    }

    // extract keys from kvs to array 
    fmt.Println(len(kvs))
    key_arr := make([]string, 0, len(kvs))
    for k := range kvs {
        key_arr = append(key_arr, k)
    }
    fmt.Println(key_arr)

    // range on strings iterates over Unicode code points
      for i, c := range "go lang" {
            fmt.Println(i, c)
        }

}