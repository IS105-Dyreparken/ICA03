package main

import (
	"fmt"
	"os"
)

var expString = "Nord og Sør"
var langString = os.Args[1]

func main() {
	Translate(expString, langString)
}

func Translate(expString string, langString string) {
	if langString == "jp" {
		printJP()
	} else if langString == "is" {
		printIS()
	}
}

//japansk
func printJP() string {
	jp := "\xE2\x80\x9C\xE5\x8C\x97\xE3\x81\xA8\xE5\x8D\x97\xE2\x80\x9D"
	fmt.Printf("%s\n", jp)
	return jp
}

//Islansk
func printIS() string {
	is := "\xE2\x80\x9C\x6E\x6F\x72\xC3\xB0\x75\x72\x20\x6F\x67\x20\x73\x75\xC3\xB0\x75\x72\xE2\x80\x9D"
	fmt.Printf("%s\n", is)
	return is
}
