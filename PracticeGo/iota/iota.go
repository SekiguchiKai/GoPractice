//iota

package main

import "fmt"

func main() {
	const (
		Go = iota + 1
		Java
		C
		Python
		PHP
		JavaScript
	)

	fmt.Println(Go, Java, C, Python, PHP, JavaScript)
}
