// Nombre de la carpeta donde estan todos los archivos
package comunida

import "fmt"

// Estoy exportando a persona1
func Casa() {
	//podemos usar la variables otros archivos sin usar import o export
	fmt.Println(persona1)
}

func Casa2() {
	//solo usando el nombre de la variable
	fmt.Println(persona2)
}
