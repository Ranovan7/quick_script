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
}
