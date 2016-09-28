//map
//mapは[key]value

package main

import "fmt"

func main() {
	//makeでの生成
	staticLan := make(map[int]string)
	staticLan[5] = "Java"
	staticLan[23] = "Go"
	staticLan[75] = "TypeScript"
	fmt.Println(staticLan)
	fmt.Println(len(staticLan))

	//リテラルでのmap生成
	modernJS := map[string]string{"FrameWork": "Angular2", "AltJS": "TypeScript", "ESver": "ECMAScript6"}
	fmt.Printf("%s", "modernJSの中身は")
	fmt.Println(modernJS)

	//"AltJS": "TypeScript"を削除
	delete(modernJS, "AltJS")
	fmt.Printf("%s", "AltJS削除後のmodernJSは")
	fmt.Println(modernJS)

	


	


}
