//Method

/*
データ型にバンドルされた関数

func (レシーバの変数名　レシーバーの型) 関数名(引数　引数の型) 戻り値の型 {
	処理
	バンドルされた構造体のフィールドを使いたい時には、構造体の入っている変数.フィールド名
}

*/

package main

import "fmt"

type kaimono struct {
	egg int
	apple int
}

func (k kaimono) add(x, y int) int{
	return k.egg * x + k.apple * y
}

func main()  {
	k :=kaimono{egg:100, apple:200}
	fmt.Printf("%s", "お買い物の合計は、" )
	fmt.Print(k.add(5, 7))
	fmt.Printf("%s\n", "円" )
}

