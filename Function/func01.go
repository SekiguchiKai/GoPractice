
//goの関数その1
package main
import (
    "fmt"

)

func taxCalc(a, b, c float64) float64 {
    tax :=1.08
    return (a + b + c) * tax
}


//複数の引数を返す
func plusMultiple (e, f int) (int, int) {
    plus := e + f
    multiple := e * f
    return plus, multiple
}

//無名関数
triangle := func(height, base int) int {
    area := height * base
    return area
} 

func main() {
    fmt.Println(taxCalc(100.0, 200.0, 300.0))
    fmt.Println(plusMultiple(100, 200))
    fmt.Println(triangle(50, 20))
}