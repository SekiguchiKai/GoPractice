
package main

import "fmt"


type Saiyazin struct {
	name string
	hp int
}

//値渡し
func (Goku Saiyazin) Kamehameha(rival Saiyazin) {
	rival.hp--
}

//参照渡し
func (Goku *Saiyazin) GenkiDama(rival Saiyazin) {
	rival.hp--
}

func (Bezita Saiyazin) Printout() {
	fmt.Printf("%s%d\n",  "ベージタのHPは", Bezita.hp)
}

func (Nappa Saiyazin) Printout2() {
	fmt.Printf("%s%d\n", "ナッパのHPは", Nappa.hp)
}

func main()  {

	Goku := Saiyazin{name:"悟空", hp:20}


	Bezita := Saiyazin{name:"ベジータ", hp:20}

	Nappa := Saiyazin{name:"ナッパ", hp:10}


	Bezita.Printout()
	fmt.Println("悟空かめはめ波三連ちゃん")
	Goku.Kamehameha(Bezita)
	Goku.Kamehameha(Bezita)
	Goku.Kamehameha(Bezita)
	Bezita.Printout()


	Bezita.Printout2()
	fmt.Println("悟空元気玉三連ちゃん")
	Goku.GenkiDama(Nappa)
	Goku.GenkiDama(Nappa)
	Goku.GenkiDama(Nappa)
	Nappa.Printout2()


}


/*以下は、エラーになる
→構造体をnewでポイント型生成したため
→newでは指定したがたのポインタ型を生成する
→ポインタは、メモリ領域のアドレスを指し示しているのでnewするたびに、Saiyazin型の
　あるメモリのアドレスが、変わってしまうから

package main

import "fmt"


type Saiyazin struct {
	name string
	hp int
}

//値渡し
func (Goku Saiyazin) Kamehameha(rival Saiyazin) {
	rival.hp--
}

//参照渡し
func (Goku *Saiyazin) GenkiDama(rival Saiyazin) {
	rival.hp--
}


func main()  {

	Goku := new(Saiyazin)
	Goku.name = "悟空"
	Goku.hp = 20

	Bezita := new(Saiyazin)
	Bezita.name = "ベジータ"
	Bezita.hp = 20

	Nappa := new(Saiyazin)
	Bezita.name = "ナッパ"
	Bezita.hp = 10


	fmt.Println("悟空かめはめ波三連ちゃん")
	Goku.Kamehameha(Bezita)
	Goku.Kamehameha(Bezita)
	Goku.Kamehameha(Bezita)
	fmt.Printf("%s", "かめはめ波後のベージタのHPは")
	fmt.Println(Bezita.hp)

	fmt.Println("悟空元気玉三連ちゃん")
	Goku.GenkiDama(Nappa)
	Goku.GenkiDama(Nappa)
	Goku.GenkiDama(Nappa)
	fmt.Printf("%s", "元気玉後のナッパのHPは")
	fmt.Println(Bezita.hp)

}

*/