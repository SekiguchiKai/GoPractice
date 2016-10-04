//構造体

/*複数の型の値を意味のある新しい型として定義することができる
構造体の中に構造体を含めることも可能
構造体と手続きを結びつけるmethod
*/

//メソッド＝＞データ型にバンドルされた関数
/*データ型の外にfuncを定義する*/

/*基本構文
type 構造体名 struct {
    この構造体型に含めたいデータを書いていく
    このデータが各々フィールド呼ばれる
}

変数 := new(構造体名)で、構造体分の領域をメモリに確保、初期化をしてポインタを返す

*/
package main

import "fmt"

type person struct {
	name string
	age int
}


func main() {
	//personを生成して、myに代入
	my := new(person)
	//meのデータを埋めていく
	my.name = "sekky"
	my.age = 26
	fmt.Println(my)

	//リテラルで生成
	my2 := person{"sekky2", 27}
	fmt.Println(my2)
}

