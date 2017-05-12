package main

import "fmt"

func main() {
	a := "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"
	fmt.Println("Med %s")
	fmt.Printf("%s\n", a)
	fmt.Println("Med %q")
	fmt.Printf("%q\n", a)
	fmt.Println("Med %+q")
	fmt.Printf("%+q\n", a)
	fmt.Println("Med % X")
	fmt.Printf("% X\n", a)
	fmt.Println("Med %c")
	fmt.Printf("%c\n", a)
}
