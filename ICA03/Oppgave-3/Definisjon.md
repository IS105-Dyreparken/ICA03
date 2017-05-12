# Oppgave 3

%s
Interpreterer ikke stringen, altså returnerer den akkurat slik den er skrevet inn.
```
func main() {
	a := "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"
	fmt.Printf("%s", b)
}
    returnerer ��=� ⌘ //Avhenger av hvilket format teksteditoren din har.
```

%q
Interpreterer en tekststreng og returnerer alt som er et unicode code point. De delene av strengen som ikke tilsvarer et unicode code point blir skrevet som ren kode.

```
func main() {
	a := "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"
	fmt.Printf("%q", b)
}
    returnerer "\xbd\xb2=\xbc ⌘"
    Avhenger igjen hvilken teksteditor du har, men her kan vi set at den returnerer ren kode for de tegnene den ikke klarte å tolke.
```

%+q
```
func main() {
	a := "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"
	fmt.Printf("%q", b)
}
    returnerer "\xbd\xb2=\xbc \u2318"
    Her kan vi set at den også har tolket om Unicode tegned til UTF-8 kode.
```



%X
Skriver ut de heksadesimale verdiene til den strengen man bruker som input.
```
func main() {
	a := "Jesus"
	fmt.Printf("%q", b)
}
    returnerer 4A65737573
```

% X
Samme som %X, bare at den har spaces i mellom de forskjellige verdiene.
```
func main() {
	a := "Jesus"
	fmt.Printf("%q", b)
}
    //returnerer 4A 65 73 75 73
    Her kan vi se at den har "oversatt" strengen vår til Base16, eller heksadesimale tegn.
```

%c
Skriver ut en og en karakter om gangen og prøver å tolke de ut ifra ASCII(extended) kodesettet
```
func main() {
	a := "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"
	fmt.Printf("%c", b)
}
    returnerer
```
