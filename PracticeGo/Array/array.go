//Array

package main

import "fmt"

func main() {
	//一般的な配列の宣言
	var staticLang [3]string
	staticLang[0] = "Java"
	staticLang[1] = "Go"
	staticLang[2] = "C"
	fmt.Println(staticLang)
	//要素の個数を表示
	fmt.Printf("要素の個数は")
	fmt.Println(len(staticLang))

	dynamicLang := [3]string{"Python", "Ruby", "PHP"}
	fmt.Println(dynamicLang)

}
