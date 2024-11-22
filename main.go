package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type Alumno struct {
	Nombre string
	Notas  []float64
}

func leerCsv(nombreArchivo string) ([]Alumno, error) {

	archivo, err := os.Open(nombreArchivo)
	if err != nil {
		return nil, err
	}
	lector := csv.NewReader(archivo)
	lector.Comma = ';'

	registros, err := lector.ReadAll()
	if err != nil {
		return nil, err
	}
	var alumnos []Alumno

	for _, registro := range registros {
		nombre := registro[0]
		var notas []float64
		for _, notasStr := range registro[1:] {
			nota, err := strconv.ParseFloat(notasStr, 64)
			if err != nil {
				return nil, err
			}
			notas = append(notas, nota)
		}
		alumnos = append(alumnos, Alumno{Nombre: nombre, Notas: notas})
	}
	return alumnos, nil

}

func main() {
	alumnos, err := leerCsv("notas.csv")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for _, alumno := range alumnos {
		fmt.Printf("Alumno: %s, Notas: %v\n", alumno.Nombre, alumno.Notas)
	}
	// var nombrePersona string = "Matias"
	// var apellido = "marensi"
	// segundonombre := "santiago"
	// fmt.Println(nombrePersona + " holaaaaa " + apellido + " " + segundonombre)

	// fmt.Println("hola mundo")

	// var año int16 = 2024
	// edad := 37
	// apellido = "gola"

	// fmt.Println(año, edad, segundonombre, apellido)

	// var frutas = [3]string{"pera", "naranja", "algo"}

	// fmt.Println(frutas[0])

	// var paises = [4]string{"arg", "chile", "bra"}

	// fmt.Println(paises)
	// paises[3] = "nuevo"
	// fmt.Println(paises)
}
