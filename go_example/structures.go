package main

import (
    "fmt"
)

type Example struct {
    Ark     string
    Bork    []int
}

func main() {
    ex := Example{
        Ark: "Arknights",
    }

    if ex.Bork == nil {
        fmt.Println("Bork is empty")
    }
}
