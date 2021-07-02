package main

import "fmt"


func main () {


    // one liner declarion with 0 size
    // var intSlice []int
    // var strSlice []string

    // var intSlice = make([]int, 10)        // when length and capacity is same
    // var strSlice = make([]string, 10, 20) // when length and capacity is different

    s := make([]string, 60)
    fmt.Println("emp:", s)

    s[0] = "a"
    s[1] = "b"
    s[2] = "2"
    s[8] = "d"
    s[59] = "X"
    fmt.Println("set:", s)
    fmt.Println("get:", s[2])

    // print the length
    fmt.Println("length",len(s))

    //dynamic pushing data in slices
    //this will increase the size
    s = append(s, "d")
    s = append(s, "e", "f")
    fmt.Println("apd:", s)

    // size increased
    fmt.Println("length > ",len(s))


    // clone slice 
    c := make([]string, len(s) + 10 )

    // this holds the reference
    c = s
    print("\n\n")
    fmt.Println(len(c),len(s))
    print("\n\n")
    fmt.Println("c before => ", c)
    fmt.Println("s before => ", s)

    // insert element in s can be seen in c as well
    s[20] = "added"
    fmt.Println("c after => ",c)
    fmt.Println("s after => ",s)

    // works the same as above
    copy(c, s)
    s[25] = "added after copy"
    fmt.Println("cpy:", c)
    // assign element
    l := s[0:len(s)]
    fmt.Println("sl1:", l)

    // print till index 5
     l = s[:5]
    fmt.Println("sl2:", l)

    // print from index 2 till end
     l = s[2:]
    fmt.Println("sl3:", l)

    // online declaration of t
    t := []string{"a","b","c","d"}
    fmt.Println("dcl:", t)
    fmt.Println("length", len(t))

    // push in slice
    x := make([]int, 0)
    fmt.Println(x)

    //slice append
    for i:= 1; i <= 1000; i++ {
        x = append(x,i)
    }

    fmt.Println("Inserted element in x =>", len(x))


    // two dimentional slice
    TwoDSlice := make([][]int, 10)


    for i:= 1 ; i < 10 ; i++ {
        // increase inner length
        innerLen := i + 1
        // set inner slice length
        TwoDSlice[i] = make([]int, innerLen)
        for j:= 1; j < innerLen; j++ {
            TwoDSlice[i][j] = i + j
        } 
    }
    fmt.Println("2d: ", TwoDSlice)

}