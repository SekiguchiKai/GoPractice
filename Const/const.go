//定数の宣言

/*
const 定数名 = 定数
*/

package main

import (
	"fmt"
	"math"
)

func okaimono(apquantity, baquantity int) float64 {
	const tax = 1.08
	apple := 100
	banana := 200

	amount := (float64)(apple*apquantity+banana*baquantity) * tax
	return amount
}

func main() {
	result := math.Trunc(okaimono(2, 3))
	fmt.Printf("お買い物の合計金額は %f円です\n", result)
}
