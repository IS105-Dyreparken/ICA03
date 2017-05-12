Oppgave 1

a) For oss ser vi ikke symbol for 0 til 1f og 7F. Det samme skjer på alle platformer vi har testet. (Mac, Windows og i nettskyen)

b)Se ascii.go

ascii_main.go importerer alle pakkene som ligger inn i mappen ./ascii Her inne ligger pakkene ascii.go og ascii_test.go.

Funksjonen genererer en utskrift fra en sekvens av bytes, dvs. av typen []bytes Funksjonen greetingASCII() returnerer en variabel av typen string, som inneholder kun ASCII-tegn.

c)Vi har implementert en test for funksjonen GreetingASCII i en egen fil ascii_test.go, som tester om returverdier av typen string inneholder kun ASCII-tegn.

Alle utførelsene av ascii_main.go får samme resulateter på hver platform vi tester det på.
