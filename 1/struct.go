package main

import "fmt"

type Vertex struct {
	X int
	Y int
	S string
}

func main() {
	v := Vertex{X: 1, Y: 2}
	fmt.Println(v)
	fmt.Println(v.X, v.Y)

	xx := Vertex{1, 2, "test"}
	fmt.Println(xx)

	yy := Vertex{}
	fmt.Println(yy)

	v1 := new(Vertex)
	fmt.Println(v1)
	fmt.Printf("%T\n", v1)

	// 構造体のポインタを取得
	v2 := &Vertex{}
	fmt.Println(v2)
	fmt.Printf("%T\n", v2)

	s := make([]int, 10)
	fmt.Println(s)
	fmt.Printf("%T\n", s)

	fmt.Println("---------")
	x1 := Vertex{X: 1, Y: 2}
	changeVertex(x1)
	fmt.Println(x1)

	x2 := &Vertex{X: 1, Y: 2}
	changeVertex2(x2)
	fmt.Println(x2)
}

// 値渡し
func changeVertex(v Vertex) {
	v.X = 1000
}

func changeVertex2(v *Vertex) {
	v.X = 1000
	(*v).Y = 2000
}
