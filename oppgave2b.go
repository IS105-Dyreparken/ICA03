package main

import "fmt"

func main() {
	//
	var greeting []byte
	for i := 0x80; i <= 0xFF; i++ {
		greeting = append(greeting, byte(i))

		greeting := "\x53\x61\x6C\x75\x74\x2C\x20\xE7\x61\x20\x76\x61\xB0\x2D\x29\x80\x35\x30"
		fmt.Println(string(greeting))

	}
}
