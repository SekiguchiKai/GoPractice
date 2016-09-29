//Pointer
//変数を宣言するとメモリ上に領域（アドレス）を確保する
//pointerは、そのアドレスを指し示す変数
//演算は、できない

package main

import "fmt"

func main() {
	/*変数xを宣言し、数値を代入
	  　これによって、メモリ上にx用の領域（アドレス）が確保される
	*/
	x := 100

	//その領域を確保するための変数であるpointerを宣言
	var xPo *int
	xPo = &x //xPo = xのアドレス

	fmt.Println(xPo)

	//xPoの/領域にあるデータの値 => *xPo
	fmt.Println(*xPo)

}
