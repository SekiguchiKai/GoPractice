//range
/*配列、slice、mapと組み合わせて使う
要素分だけ何らかの処理を繰り返したい時につかう
*/

/*for i, v := range 変数名 {
繰り返し処理
}

i = sliceのインデックス（何番目の要素か）
v = sliceの値
*/

/*
iかvのどちらかが必要ない時に、必要のないi（もしくはv）を繰り返し処理の中で
破棄する方法
for i, v := range 変数名 {
    繰り返し処理
}
のiもしくはvの必要のない方を"_"（アンダーバーにする）
繰り返し処理の中では必要のないiもしくはvは書かない
*/

package main

import "fmt"

func main() {
	staticLan := []string{"Java", "Go", "C"}
	for i, v := range staticLan {
		fmt.Println(i, v)
	}

	//vだけを表示
	for _, v := range staticLan {
		fmt.Println(v)
	}
	//mapでrange
	webTech := map[string]string{"FrontSide": "JavaSctipt", "ServerSide": "Go", "DataBase": "MySQL"}
	for k, v := range webTech {
		fmt.Println(k, v)
	}

}
