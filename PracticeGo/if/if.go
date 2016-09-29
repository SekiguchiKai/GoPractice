//if 

package main

import"fmt"

func main()  {

    Pro := "Go"
    if Pro == "Java" || Pro == "Go" || Pro == "TypeScript" {
        fmt.Printf("%s\n", "静的言語です")
    } else if Pro == "Ruby" || Pro == "JavaScript" || Pro == "Python" {
        fmt.Printf("%s\n", "動的言語です")
    } else {
        fmt.Printf("%s\n", "I do not know")
    }
    
}