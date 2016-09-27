//Slice
/*
Goの配列は、参照型ではない
配列の要素数は、固定長
sliceは参照型
sliceは、可変長
*/

package main

import "fmt"

func main() {
	//組み込み関数makeを利用したスライスの生成
	s := make([]int, 3)
	fmt.Println(s)

	s2 := make([]string, 3)
	s2[0] = "JavaScript"
	s2[1] = "TypeScript"
	s2[2] = "ECMAScdript"
	fmt.Println(s2)

	//スライスの要素数を調べる
	fmt.Println(len(s2))

	//リテラルでのスライスの生成
	s3 := []string{"Swift", "Scala", "R"}
	fmt.Println(s3)
	fmt.Printf("%s", "s3の要素数は、")
	fmt.Println(len(s3))

	s4 := []string{"C"}
	fmt.Println(s4)

	//"C++", "C#", "Objective-C"をスライスの末尾に追加
	s4 = append(s4, "C++", "C#", "Objective-C")
	fmt.Println(s4)

}
