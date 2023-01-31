package main

import "fmt"

func main() {
	x := 100
	one(x)
	println(x)

	// xのアドレスを渡すため値が書き換わる
	two(&x)
	println(x)

	var p *int = new(int) // メモリ領域を確保している
	println(p)
	println(*p) // 0
	*p++
	println(*p) // 1

	var p2 *int // メモリ領域を確保していない
	println(p2)
	fmt.Println("---------")

	// newとmakeの違い
	// newはポインタを返す
	s := make([]int, 10) // スライスの作成
	fmt.Printf("%T\n", s)

	m := make(map[string]int) // マップの作成
	fmt.Printf("%T\n", m)

	ch := make(chan int) // チャネルの作成
	fmt.Printf("%T\n", ch)

	i := new(int) // ポインタの作成
	fmt.Printf("%T\n", i)

	st := new(struct{}) // 構造体の作成
	fmt.Printf("%T\n", st)
}

func one(x int) {
	x = 1
}

/*
ポインタを引き受ける
*/
func two(x *int) {
	*x = 2
}
