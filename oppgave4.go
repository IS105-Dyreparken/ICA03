package main

import "fmt"

func main() {
	const placeOfInterest = "\xbd\xb2\x3d\xbc"

	fmt.Printf("plain string: ")
	fmt.Printf("%s", placeOfInterest)
	fmt.Printf("\n")

	fmt.Printf("quoted string: ")
	fmt.Printf("%+q", placeOfInterest)
	fmt.Printf("\n")

	fmt.Printf("hex bytes: ")
	for i := 0; i < len(placeOfInterest); i++ {
		fmt.Printf("% X ", placeOfInterest[i])
	}

	fmt.Printf("\n")
	alt()
}

func alt() {

	const oppgave3 = "\xBD\xB2\x3D\xBC\x20\xe2\x8c\x98"

	fmt.Printf("string: ")
	//for i := 0; i < len(oppgave3); i++ {
	//fmt.Printf("%s", oppgave3[i])
	fmt.Printf("%s\n", oppgave3)
}
