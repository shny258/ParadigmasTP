package main

import (
	"fmt"
	"gestorEstudiantesEnGo/Estudiante"
	"time"
)

func main() {
	e1 := Estudiante.NuevoEstudiante(123, "Lucía", time.Date(2000, 5, 10, 0, 0, 0, 0, time.UTC))
	e1.Notas = []float64{8, 9, 7}

	e2 := Estudiante.NuevoEstudiante(456, "Mateo", time.Date(2001, 8, 20, 0, 0, 0, 0, time.UTC))
	e2.Notas = []float64{5, 6, 6}

	gestor := Estudiante.NuevoGestorConSlice([]*Estudiante.Estudiante{e1, e2})

	fmt.Println("Estudiante con DNI 123:", gestor.BuscarEstudiantePorDNI(123))
	fmt.Println("Aprobados:")
	for _, e := range gestor.ListarEstudiantesPromedioAprobado() {
		fmt.Println(e.Nombre)
	}
}
