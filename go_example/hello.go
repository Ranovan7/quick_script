package main

import "fmt"

func main() {
    fmt.Println("Hello World")

    a := []int{21, 22, 23}
    fmt.Println(a)

    var i int
    i = 4/(4/2)
    fmt.Println(i)

    i_arr := []int{1,2,3,4}
    min := i_arr[0]
    for i := 1; i < len(i_arr); i++ {
        if min > i_arr[i] {
            min = i_arr[i]
        }
    }
    fmt.Println(min)

    var x *int
    fmt.Println("x = ", x)

    y := 10
    x = &y
    fmt.Println(*x)
}
