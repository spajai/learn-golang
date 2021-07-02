package main

import f "fmt"
// to import time 
import t "time"

func main () {

    // interger i = 2
    i:= 2
    // switch else go to default
    switch i {
        case 1:
            f.Println("You have selected ", i)
        case 2:
            f.Println("You have selected ", i)
        default:
            f.Println("You have selected none")
    }
    
    // this will print todays day day of week
    f.Println(t.Now().Weekday())

    // this will literally print the day
    f.Println(t.Saturday)
    
    switch t.Now().Weekday() {
    case t.Saturday, t.Sunday:
        f.Println("It's the weekend")
    default:
        f.Println("It's a weekday")
    }


    // it will get the time
   t_now := t.Now()
    switch {
    case t_now.Hour() < 12:
        f.Println("It's before noon")
    default:
        f.Println("It's after noon")
    }


    // function def
    WhatTypeIam := func (i interface {}) {
        // obtained type of i
        switch t := i.(type) {
            case bool:
                f.Println("I am bool")
            case int: 
                f.Println("I am int")
            default:
            // printf %T will print the type
                f.Printf("I dont know the type %T \n", t)
        }
    }

    // call the fun
    WhatTypeIam(true)
    WhatTypeIam(1)
    WhatTypeIam("Test Str")
    

}