//goのswitch文と擬似乱数

package main 

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    programLang := [5] string {"Go", "Java", "Python", "ruby"} 
    rand.Seed(time.Now().UnixNano())

    switch programLang[rand.Intn(5)] {
    case "Go", "Java":
        fmt.Println(programLang[rand.Intn(5)] + " is static programming language")
    case "Python", "ruby":
        fmt.Println(programLang[rand.Intn(5)] + " is dynamic programming language")
    default:
        fmt.Println("unknown")
    }
}