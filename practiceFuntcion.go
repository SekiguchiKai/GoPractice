package main 

import "fmt"

func triangleArea(a int, b int) int {
    return a * b /2
}

func main() {
    fmt.Println(triangleArea(10,5))
}