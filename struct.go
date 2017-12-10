package main

import (
    "fmt"
    "math"
)

type Point struct {
    x float64
    y float64
}

func norm(a *Point, b *Point) float64 {
    return math.Sqrt((a.x - b.x)*(a.x - b.x) + (a.y - b.y)*(a.y - b.y))
}

func main() {
    p := Point{1.2, 2.5}
    l := Point{4.2, 5.5}

    fmt.Println(norm(&p, &l))
}
