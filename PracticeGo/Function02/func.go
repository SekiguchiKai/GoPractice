//即時実行関数

package main

import "fmt"

func main() {
	func(greeting string) {
		fmt.Printf("%s\n", greeting)
	}("Hello,Go!!!!!!!")
}
