//構造体

/*複数の型の値を意味のある新しい型として定義することができる
構造体の中に構造体を含めることも可能
構造体と手続きを結びつけるmethod
*/

//メソッド＝＞データ型にバンドルされた関数
/*データ型の外にfuncを定義する*/
package main 

import "fmt"

type System struct {
    serverSide string
    frontEnd string 
    dataBase string 
}

func main()  {
    //通常のnew
    s1 := new(System)
    s1.serverSide = "Go"
    s1.frontEnd = "JavaScript"
    s1.dataBase = "MySQL"
    fmt.Println(s1)

    //値を与えながらnew
    s2 := System{"Java","frontEnd","Oracle"}
     fmt.Println(s2)
}

