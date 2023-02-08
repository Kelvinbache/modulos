package archivos

type usuarios struct {
	Nombre   string
	Apellido string
	Edad     int
	Telefono int
}

func datos(nombre, apellido string, edad, telefono int) {
	usuario := []usuarios{
		{Nombre: nombre, Apellido: apellido, Edad: edad, Telefono: telefono},
	}

}
