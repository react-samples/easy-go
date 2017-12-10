package main

import "fmt"


func main(){
    var a []int = []int{1, 2, 3, 4, 5, 6}
    a = append(a, 7)
    for _, i := range a {
	      fmt.Println(i)
    }
}
