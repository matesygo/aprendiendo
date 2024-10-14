package main

import "testing"

func TestSaludo(t *testing.T) {
	tengo := Saludo("Chris")
	quiero := "Hola, Chris"

	if tengo != quiero {
		t.Errorf("tengo %q quiero %q", tengo, quiero)
	}
}
