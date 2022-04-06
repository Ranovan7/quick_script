package main

import (
    "crypto/md5"
    "fmt"
)

func main () {
    text := "dj91238eu102iehc17029%Bexu10928rc71023"
    hash := md5.Sum([]byte(text))

    fmt.Println(text)
    fmt.Println(hash)
    fmt.Println(fmt.Sprintf("%x", hash))
}
