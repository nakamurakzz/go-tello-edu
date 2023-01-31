package main

import (
	"fmt"
	"os"
)

func main() {
	file, _ := os.Open("1/defer.go")
	defer file.Close() // ファイルのcloseし忘れを防ぐ

	// ファイルの内容を読み込む
	data := make([]byte, 200)
	file.Read(data)
	fmt.Println(string(data))
}
