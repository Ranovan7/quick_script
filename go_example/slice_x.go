package main

import (
    "fmt"
    "strings"
)

func main() {
    deleted := []string{}

    deleted = append(deleted, "345662")
    deleted = append(deleted, "654321")
    deleted = append(deleted, "973554")

    fmt.Println(deleted[0])

    message := fmt.Sprintf("Deleted IDs : (%s)", strings.Join(deleted, ", "))
    fmt.Println(message)

    example := []int{1,2,3,4,5}
    sum := 0
    for _, example := range(example) {
        sum += example
    }
    fmt.Println(sum)
    fmt.Println(example)

    a := []int{1,2,3,4,5}
    b := []int{7,8,9}
    a = append(a, b...)
    fmt.Println(a)
}
