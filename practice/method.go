package main

type Vertex struct {
	x int
	y int
}

func (v Vertex) Area() int {
	return v.x * v.y
}

/*
インスタンスの値を変更する
*/
func (v *Vertex) scale(i int) {
	v.x = v.x * i
	v.y = v.y * i
}

/*
コンストラクタ、structを返すGoのデザインパターン
*/
func New(x, y int) *Vertex {
	return &Vertex{x, y}
}

func main() {
	v := Vertex{x: 10, y: 20}
	println(v.Area())
	v.scale(10)
	println(v.Area())

	v2 := New(1, 2)
	println(v2.Area())
}
