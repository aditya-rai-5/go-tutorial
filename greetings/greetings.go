// package main
package greetings

import (
	"fmt"
)

func Hello(name string) string {
	message := fmt.Sprint("hii", name)
	return message
}

// func main() {
// 	out := Hello("aditya")
// 	fmt.Println(out)
// }
