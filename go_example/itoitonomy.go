package main

import "fmt"

func main() {
    answer := "ABC"
    number := 10
    somestrstr := "answer is = %s, with number = %d"
    somestr, _ := fmt.Printf(somestrstr, answer, number)

    fmt.Println(somestr)

    // arr := []int{1,2,3,4,5}
    // val, isset := arr[7]
    // fmt.Println(val)
    // fmt.Println(isset)
}
