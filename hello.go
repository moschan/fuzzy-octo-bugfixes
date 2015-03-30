package main

// import "fmt"
import "hello"

// var foo string = "Hello, World!" // 型注釈

func main() {
	// foo := "Hello, World!"

	// 計算
	// result := 0;
	// result = (15 + 35) / 2

	// 配列
	// sl := make([]int, 0) //=> []
	// appendキーワードで追加
	// sl = append(sl, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0)

	// result := addRIch("Media")
	//
	// fmt.Println(result)

	hello.SayHello()
}

func addRIch(str string) (result string) {
	result = "Rich" + str
	return
}
