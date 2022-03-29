package main

import (
    "fmt"
    "strings"
)

func main() {
    somestr := "SOMEstring"
    fmt.Println(strings.ToLower(somestr))

    id := 12345
    meB := "MeBeyonce"
    ex := fmt.Sprintf("id=%d&try=%s", id, meB)
    fmt.Println(ex)

    ex = fmt.Sprintf("%s&id=%d&try=%s", ex, id, meB)
    fmt.Println(ex)

    if nil != "" {
        fmt.Println("Nil is not the same as Empty String")
    }
}
