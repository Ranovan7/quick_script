package main

import "fmt"

const PI = 3.14

func main() {
    var a int = 10
    b := 15

    fmt.Println(a)
    fmt.Println(b)

    var str_a string = "Hello World"
    var floating float32 = 0.34
    a_bool := true

    fmt.Println(str_a)
    fmt.Println(floating)
    fmt.Println(a_bool)

    var x string
    var y int
    var z bool

    fmt.Println(x)
    fmt.Println(y)
    fmt.Println(z)

    var q, w, e float32 = 0.3, 3.4, 82.9
    fmt.Println(q, w, e)

    a1, a2 := "I love you", 3000
    fmt.Print(a1)
    fmt.Print(a2)
    fmt.Print("\n")

    var (
        p string = "Pi is"
        notval float32 = 22.7
    )
    fmt.Printf("%s %.2f (%T) not %.1f (%T)\n", p, PI, PI, notval, notval)

    arr := [...]int{1,3,5,7}
    arr[3] = 9
    fmt.Println(arr)

    arr1 := [...]int{1:10,2:40}
    fmt.Println(arr1)
    fmt.Println(len(arr1))

    myslice2 := []string{"Go", "Slices", "Are", "Powerful"}
    fmt.Println(len(myslice2))
    fmt.Println(cap(myslice2))
    fmt.Println(myslice2)

    arrx := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    fmt.Println(arrx)
    fmt.Println(arrx[:3])

    is_slice := arrx[3:7]
    fmt.Printf("arrx is %T and is_slice is %T\n", arrx, is_slice)

    myslice1 := make([]int, 5, 10)
    fmt.Printf("myslice1 = %v (%T)\n", myslice1, myslice1)
    fmt.Printf("length = %d\n", len(myslice1))
    fmt.Printf("capacity = %d\n", cap(myslice1))

    myslice1 = append(myslice1, 11, 12)
    myslice1[0] = 1
    fmt.Println(myslice1)
    fmt.Println(myslice1[5:8])

    myslice1 = append(myslice1, 13, 14, 15, 17)
    fmt.Println(myslice1)
    fmt.Printf("length = %d\n", len(myslice1))
    fmt.Printf("capacity = %d\n", cap(myslice1))

    myslice3 := []int{21, 22, 23}

    myslice1 = append(myslice1, myslice3...)
    fmt.Println(myslice1)
    fmt.Printf("length = %d\n", len(myslice1))
    fmt.Printf("capacity = %d\n", cap(myslice1))
}
