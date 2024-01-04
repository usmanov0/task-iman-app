package main

import "fmt"

func main() {
	doSomeAction()
}

func doSomeAction() {
	for i := 0; i < 10; i++ {
		fmt.Println("hello world")
	}
}
