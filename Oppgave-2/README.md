# Oppgave 2


Iso_main.go importerer alle pakkene som ligger inne i ./iso. Det er Iso.go og
iso_test.go.

Iso_main.go har en func main hvor den henter ut funksjonene fra pakkene i
./iso.

func IterateOverASCIIStringLiteral(s string) formaterer det til %x, %q, %b.

" %x = Base 16, with lower-case letters for a-f
  %q = a single-quoted character literal safely escaped with Go syntax.
  %b = base 2 "
Kilde: https://golang.org/pkg/fmt/

Fra 0x80 til 0x9F viser ikke noe tegn, fordi det er kommandoer som Pc-en skal
utføre, så det er ikke noe tegn.


GreetingExtendedASCII skal skrive ut "Salut, ça va °-) €50", men her kommmer ikke
eurotegnet med. Euro-tegnet dukker ikke opp på noen av platformene vi tester det
på. Det er fordi det ligger på en annen plass i Ascii-tabellen.


I greetingExtendedASCII() er det implementert en test i egen fil iso_test.go,
som tester om returverdier (av type string) inneholder kun tegn fra en
Extended ASCII
