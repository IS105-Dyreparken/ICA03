Oppgave 3a


I denne oppgaven skal vi forklare hvordan vi gjør en omregning fra unicode til utf8 forgår.
Det vi har lært fra lesestoff som forelseren har lagt ut at Unicode rom (space) er per i dag definert fra U+0000 til U+10FFFF (hele rommet er 4 bytes = 32 bits). U+0000 - U+FFFF er “plane 0” U+10000 - U+1FFFF er “plane 1” osv. “planes” 1 - 16 (desimalt) er tilleggsrom (supplementary planes) U+100000 - U+10FFFF er siste (17-de) “plane”.


Det ½ tegnet  har kode BD I iso 8859-1 og 00BD skrives vanligvis u+00BD I unicode “plane 0” den koden 00 som kommer før BD er da “row” og den siste delen blir da kalt for column. Måten vi regner det ut på går som følgende.

Plane = 0x00 = 0b00000000
Row= 0x00 = 0b00000000
Column= 0xBD = 0b10111101

Da vil det se slik ut:

-Plane-               -Row---             Column
00000000          00000000        10111101
                                        YYY        YYZZZZZZ

Når vi skal finne finne koden I utf 8 så må vi dele opp koden til to kategorier (y og Z) slik vil det se ut.

Byte 0 =  11000010                   byte 1 =  10111101
                       YYYYY                                         ZZZZZZ

Fra byte 0 kan vi se at 1100 av koden er da C og resten er 2 så koden er da C2 I askii code.
Fra byte 1 kan vi se at 1011 av koden er B og resten er da D og til sammen blir de BD I askii code.
Etter den endringen er gjort og vi har funnet aski kodene kan vi legge det I golang og kjøre
fmt.Printf("%s", "\x22\xC2\xBD\x3F\x3D\x3F\x20\xE2\x8C\x98\x22"
for å kunne riktig “tegn” ved å kjøre printf %s metoden.


Vi har kjørt de andre formaterringene og svaret på hvilken svar vi fikk ligger på definisjon fila.
