package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        var str string = "Hello World"
        w.Write([]byte(str))
    })

    err := http.ListenAndServe(":3000", nil)

    if err != nil {
        fmt.Println(err)
    }
}
