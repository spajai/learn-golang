package main

import "fmt"

type person struct {
    name string
    age int
}


func newPerson (name string) *person {

    // values pass to struct
    p := person{name : name}
    // set the other attributre
    p.age = 42

    return &p
}

func main () {

    fmt.Println(person{"bob",20})

    fmt.Println(person{name : "alice"})

    fmt.Println(person{name : "fred", age: 60})

    fmt.Println(person{name : "fred", age: 60})

    fmt.Println(&person{name : "test",age: 40})


    // following output
    // &{test 40}
    // &{john 42}
    fmt.Println(newPerson("john"))
    s:= person{name:"test",age: 20}


    fmt.Println(s)
    fmt.Println(s.name)
    fmt.Println(s.age)

    sp:= &s
    fmt.Println(sp.age)

    // both will have this value since mutable
    sp.age = 51
    fmt.Println(sp.age)
    fmt.Println(s.age)


 }