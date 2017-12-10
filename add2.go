package main

import "fmt"

func add(a float64, b int) float64 {
    return a + float64(b)
}

func main(){
   fmt.Println(add(1.1, 2))
}
