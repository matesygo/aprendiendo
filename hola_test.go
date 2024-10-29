package main

import "testing"

func TestSaludo(t *testing.T) {
	t.Run("decirle 'Hola' a la gente", func(t *testing.T) {
		tengo := Saludo("Chris")
		quiero := "Hola, Chris"

		if tengo != quiero {
			t.Errorf("tengo %q quiero %q", tengo, quiero)
		}
	})

	t.Run("saludar al mundo entero cuando no se ingrese ningun nombre/cadena de texto como argumento", func(t *testing.T) {
		tengo := Saludo("")
		quiero := "Hola, mundo"

		if tengo != quiero {
			t.Errorf("tengo %q quiero %q", tengo, quiero)
		}
	})
}
