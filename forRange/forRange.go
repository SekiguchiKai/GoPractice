//for-range 配列ver

/*
for 配列のインデックス, 配列の要素　:= range 配列名 {
    処理
}
*/

package main

import (
    "fmt"
)

func main()  {

    //配列の生成
programLang := [5]string{"Go", "Java", "C", "python", "Swift"}

//for分
for i, p := range programLang {
    fmt.Printf("program Language(%d) is %s\n", i, p )
} 

return
    
}

