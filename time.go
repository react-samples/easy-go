package main

import (
    "fmt"
    "net/http"
    "time"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        t := time.Now()
	      format := "2006/01/02 15:04:05"
	      year := t.Format(format)
	      w.Write([]byte(year))
    })

    if err := http.ListenAndServe(":3000", nil); err != nil {
        fmt.Println(err)
    }
}
