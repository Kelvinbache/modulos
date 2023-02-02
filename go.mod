// Aqui creamos un modulo que conecta todo los demas archivos 
// comando que usamos para crear el modulo (go mod init nombre del modulo)
module caja

//esto es la version de go que tenemos
go 1.19


//Aqui estamos requiriendo un paquete de Go usando el comando go get nombre del Paquete
require github.com/alwindoss/morse v1.0.1

//nota: Para desintalar un paquete que no uses es (go mod tidy) quita el paquete que no usas  
