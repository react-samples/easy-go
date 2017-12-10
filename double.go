package main

import "fmt"

func double(a int, b int) (int, int) {
    return 2*a, 2*b
}

func main(){

    c, d := double(1, 2)
    fmt.Println(c)
    fmt.Println(d)
}
