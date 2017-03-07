package main

import (
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", foo)
	http.ListenAndServe(":3000", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
	// Her kan man bl. a. skrive data som skal sendes til nettleser
	// Neste linje kan brukes for å endre koding av sekvensen som sendes til nettleser
	// Standardinstilling er UTF-8
	//w.Header().Set("Content-Type", "text/html;charset=ISO-8859-1")

	t := time.Now()
	t2 := t.Format("02-01-2006 15:04")
	w.Write([]byte(t2))
	//w.Write([]rune(U+23F0))
	//w.Write([]byte("<font color=\"green\">Hvordan g\xe5r det, <b>\u16a6</b> ?</font><br/>"))
	//w.Write([]byte("\u16a6 - Thurs<br/>"))

}
