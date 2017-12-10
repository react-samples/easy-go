package main

import (
    "fmt"
    "time"
)

func main() {
    t := time.Now()
    format := "2006/01/02 15:04:05"
    year := t.Format(format)
    fmt.Println(year)
}
