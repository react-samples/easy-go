package main

import (
  "net/http"
  "text/template"
  "path/filepath"
  "time"
  "fmt"
)

type Comment struct {
    Date string
    Title string
}

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        res := template.Must(template.ParseFiles(filepath.Join("views", "index.html")))
        now := time.Now().Format("2006/01/02 15:04:05")
        o := Comment{ Date: now, Title: "テスト"}
        res.Execute(w, o)
    })

    if err := http.ListenAndServe(":3000", nil); err != nil {
        fmt.Println(err)
    }
}
