// nota: Tienes que tener un archivo ejecutor donde inicien todos los archivos
// el archivo main es el archivo donde todo inicia
package main

import (
	// Estamos importando los archivos y carpetas donde estan las funciones del programa
	"caja/modulos/comunida"
	"caja/modulos/morse"
	"fmt"
)

func main() {
	// from comunida pacake
	// estmos llamando las funciones y haciendo que inicien el programa
	comunida.Casa()
	comunida.Casa2()
	otraCarpeta := morse.Palbras("hola")
	fmt.Println(otraCarpeta)
}
