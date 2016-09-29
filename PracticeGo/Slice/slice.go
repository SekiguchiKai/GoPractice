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

	//配列からスライスを生成する
	//スライスは、配列への参照
	arrayX := [6]string{"Java", "Go", "C", "JavaScript", "Python", "Ruby"}
	sliceX := arrayX[0:3] //配列arrayXの[0]~[2]をsliceXに抽出
	fmt.Printf("%s", "静的言語は")
	fmt.Println(sliceX)

	fmt.Printf("%s", "スライスの要素数は")
	fmt.Println(len(sliceX))

	fmt.Printf("%s", "スライスの容量は")
	fmt.Println(cap(sliceX))

	//現在の{}"Java", "Go", "C"}のスライスに要素を加える
	fmt.Printf("%s", "現在のスライスの値は")
	fmt.Println(sliceX)

	sliceX = append(sliceX, "Lisp", "COBOL")
	fmt.Printf("%s", "appen後のスライスの値は")
	fmt.Println(sliceX)

	//copyを使う
	//copy先の変数を用意
	copyx := make([]string, len(sliceX))
	fmt.Println(copyx)
	//copyを利用する　copyは、要素数を返すので、elementx = 5になる
	elementx := copy(copyx, sliceX)
	fmt.Printf("%s", "cpoyxの要素数は")
	fmt.Println(elementx)

	fmt.Printf("%s", "copyxの中身は")
	fmt.Println(copyx)

}
