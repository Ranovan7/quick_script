package main

import (
    "fmt"
    "strings"
    "crypto/md5"
    "encoding/hex"
)

func GetMD5Hash(text string) string {
    hasher := md5.New()
    hasher.Write([]byte(text))
    return hex.EncodeToString(hasher.Sum(nil))
}

func AltMD5(text string) string {
    d := []byte(text)
	return fmt.Sprintf("%x", md5.Sum(d))
}

func main() {
    somestr := "SOMEstring"
    fmt.Println(strings.ToLower(somestr))

    id := 12345
    meB := "MeBeyonce"
    ex := fmt.Sprintf("id=%d&try=%s", id, meB)
    fmt.Println(ex)

    ex = fmt.Sprintf("%s&id=%d&try=%s", ex, id, meB)
    fmt.Println(ex)

    // if nil != "" {
    //     fmt.Println("Nil is not the same as Empty String")
    // }

    bodyStr := `{"filter":{"statusFpDateRange":{"start":1646038009531,"end":1648630009531}},"paging":{"size":10},"sorting":{}}`
    fmt.Println(GetMD5Hash(bodyStr))
    fmt.Println(AltMD5(bodyStr))

    fmt.Println("9d016e76a2ec85a8c7682fd694f83487" == AltMD5(bodyStr))
}
