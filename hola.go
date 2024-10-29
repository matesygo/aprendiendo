package main

const prefijoEnEspañol = "Hola, "

func Saludo(nombre string) string {
	if nombre == "" {
		nombre = "mundo"
	}
	return prefijoEnEspañol + nombre
}
