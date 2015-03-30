// src/hello/say_hello.go
package hello

import "fmt"

// 大文字なのでエクスポートされる（=外部から利用可能）
func SayHello() {
    say("Hello, World!")
}

// 小文字なのでエクスポートされない（Private）
func say(message string) {
    fmt.Println(message)
}
