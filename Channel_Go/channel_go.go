package main

import (
	"fmt"
	"time"
)

func main() {
	line := "=========="
	fmt.Println(line + "channelを使用しないgorutine" + line)
	//callerA()

	fmt.Println(line + "channelを使用したgorutine" + line)
	callerB()



}
// channelを使用しないgorutine
func callerA() {
	i := 0
	for {
		go sub1()
		sub2()

		if i > 10 {
			break
		}
		i++
	}
}
//
func callerB()  {
	// makeを使用してchannelを使用する
	// 第二引数を指定しない場合は、バッファは0
	ch := make(chan int, 10)

	go sub3(ch)

	i := 0
	for i < 10 {
		ch <- i
		i ++
	}

	time.Sleep(3 * time.Second)



}


func sub1() {
	fmt.Println("sub1 is called.")
}

func sub2() {
	fmt.Println("sub2 is called.")
}

func sub3(ch <-chan int) {
	fmt.Println("sub3 is called.")
	for {
	i := <-ch
	fmt.Printf("sub3 receves the data%v\n", i)
	}
}

func sub4(ch <-chan int) {
	fmt.Printf("sub4 receves the data%v\n", ch)
}

