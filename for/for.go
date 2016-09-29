//for.go

/*Go言語での繰り返し処理は、for文のみで、while文は存在しない
Goのfor文章は、条件式の部分に()をつかない
*/

package main

import "fmt"

func main() {

	fmt.Println("iのループの結果は")
	for i := 0; i <= 5; i++ {
		fmt.Println(i)
	}

	/*
	   While文をfor文で再現
	*/
	fmt.Println("jのループの結果は")
	j := 0
	for j < 5 {
		fmt.Println(j)
		j++
	}

	/*すべての条件を省略すると無限ループ*/

	k := 0
	fmt.Println("kのループの結果は")
	for {
		fmt.Println(k)
		k++
		if k == 10 {
			break
		}
	}

	/*break文
	  break文以降の繰り返し処理が行われなくなる（ループを抜ける）
	*/

	//5以降の処理が行われなくなる
	fmt.Println("lのループの結果は")
	for l := 0; l <= 7; l++ {
		if l == 5 {
			break
		}
		fmt.Println(l)
	}

	/*
	   continue文
	   残り処理をスキップして次の繰り返しを始める
	*/

	//3の倍数だけ表示
	fmt.Println("mのループの結果は")
	for m := 1; m <= 10; m++ {
		if m%3 != 0 {
			continue
		}
		fmt.Println(m)
	}

}
