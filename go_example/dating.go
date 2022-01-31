package main

import (
    "fmt"
    "time"
    "strings"
)

type Example struct {
    ID   int
    Qty  *int
}

func timeStrToUTC(timeStr string) string {
    tm, _ := time.Parse(time.RFC3339, timeStr)
    // return tm.UTC().Format(time.RFC3339)
    return tm.Format(time.RFC3339)
}

func main() {
    fmt.Println("Dating")

    startDate := "2021-12-11T00:00:00+07:00"
    endDate := "2021-12-12T00:00:00Z"

    startDate = strings.Replace(startDate, "Z", "+07:00", 1)
    endDate = strings.Replace(endDate, "00:00:00", "23:59:59", 1)
    endDate = strings.Replace(endDate, "Z", "+07:00", 1)

    fmt.Println(startDate)
    fmt.Println(endDate)

    fmt.Printf("Start Date : '%s'\n", timeStrToUTC(startDate))
    fmt.Printf("End Date : '%s'\n", timeStrToUTC(endDate))

    id := 10
    qty := 100
    ex := Example{ID: id, Qty: &qty}
    fmt.Println(*ex.Qty)

    OILIds := []int{549627, 549628, 549629}
    for i, _ := range OILIds {
        fmt.Println(i, OILIds[i])
    }

    example := "2022-01-31 00:00:00"
    exT := time.Parse(example)
    fmt.Println(exT)
}
