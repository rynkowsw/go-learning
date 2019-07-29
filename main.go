package main

import "fmt"

func main() {

	bar()
	fmt.Println("hello world")

	for i := 0; i < 100; i++ {

		if i%2 == 0 {
			fmt.Println(i)
		}

		fmt.Println("test2")
	}
	fu()

}
func bar() {
	fmt.Println("Let's start")
}
func fu() {
	fmt.Println("and now we exit program")
}
