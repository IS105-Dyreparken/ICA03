package main

import "fmt"

func main() {
	a := "\x22\x53\x61\x6c\x75\x74\x2c\x20\xe7\x61\x20\x76\x61\x20\xb0\x2d\x29\x20\x80\x35\x30\x22"
	b := []byte(a)
	for i := 0; i < len(b); i++ {
		fmt.Printf("%c", b[i]) // får ikke printet ut euro-tegnet
	}

}
