package main

import (
    "fmt"
)

type Example struct {
    Ark     string
    Bork    []int
    Clork   *string
}

func main() {
    // strExample := "testing"
    ex := Example{
        Ark: "Arknights",
        // Clork: &strExample,
    }

    if ex.Bork == nil {
        fmt.Println("Bork is empty")
    }

    if ex.Clork != nil {
        fmt.Println(*ex.Clork)
    }
}
