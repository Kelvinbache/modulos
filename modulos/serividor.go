package archivos

import (
	"errors"
	"os"
)

/*
pasos antes de empezar a escribir codigo:
1)primero crear los archivo donde estaran los datos del usuario
2)crear la capete donde estara el paquete de colores

pasos para el codigo del archivo servidor :
1) Primero creae la funcion y ponele parametro para los nombres
2) despues hacer una compracion entre los archivos
3) depues hacer una variable donde covirtamos los datos en json o texto


pasos para el codigo del archivo usuarios:
2) pasar directamente los datos midienate una funcion

bonificaion:
mostrar los datos del json en consola

*/

func servidor(nombreDelArchivo string){  
	archivo,error := os.Create(nombreDelArchivo)
    
	if error != nil{
	   rojo.Printf("error")
	}
  
	//conparacion entre los archivos
  _,e := os.Stat(nombreDelArchivo)	
 
  if errors.Is(e, os.ErrNotExist){

 } else if equalFile.compareFile

}