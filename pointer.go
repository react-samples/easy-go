package main

import "fmt"

func increment(a int) {
    a++
    fmt.Println(a) // 2
}

func increment_ptr(a *int){
    *a++
    fmt.Println(*a)
}

func main(){
    a, b := 1, 1

    increment(a)
    fmt.Println(a) // 1
    increment_ptr(&b)
    fmt.Println(b)
}
