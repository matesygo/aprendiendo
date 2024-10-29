package main

import "testing"

func TestSaludo(t *testing.T) {
	t.Run("decirle 'Hola' a la gente", func(t *testing.T) {
		tengo := Saludo("Chris")
		quiero := "Hola, Chris"

		evaluarMensajeCorrecto(t, tengo, quiero)
	})

	t.Run("saludar al mundo entero cuando no se ingrese ningun nombre/cadena de texto como argumento", func(t *testing.T) {
		tengo := Saludo("")
		quiero := "Hola, mundo"

		evaluarMensajeCorrecto(t, tengo, quiero)
	})
}

func evaluarMensajeCorrecto(t testing.TB, tengo, quiero string) {
	t.Helper()
	if tengo != quiero {
		t.Errorf("obtuve %q quiero %q", tengo, quiero)
	}
}
