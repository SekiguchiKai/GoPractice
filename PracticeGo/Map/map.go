//map

package main

import "fmt"

func main() {
	//makeでの生成
	staticLan := make(map[int]string)
	staticLan[5] = "Java"
	staticLan[23] = "Go"
	staticLan[75] = "TypeScript"
	fmt.Println(staticLan)

	//リテラルでのmap生成
	modernJS := map[string]string{"FrameWork": "Angular2", "AltJS": "TypeScript", "ESver": "ECMAScript6"}
	fmt.Println(modernJS)
}
