package main

import "fmt"

func main() {

	var extascii []byte
	fmt.Println(extascii)
	for i := 0x80; i <= 0xFF; i++ {
		extascii = append(extascii, byte(i))
	}
	fmt.Printf("%X ", extascii)
	fmt.Printf("%c ", extascii)
	fmt.Printf("%b\n", extascii)

}
