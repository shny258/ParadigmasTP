package estudiante

import (
	"errors"
	"time"
)

// Definimos el struct Estudiante
type Estudiante struct {
	DNI             int
	Nombre          string
	FechaNacimiento time.Time
	Notas           []float64
}

// Constructor (función que devuelve un Estudiante)
// No es realmente necesario, pero lo dejamos para mantener la misma estructura
func NuevoEstudiante(dni int, nombre string, fechaNacimiento time.Time) *Estudiante {
	return &Estudiante{
		DNI:             dni,
		Nombre:          nombre,
		FechaNacimiento: fechaNacimiento,
		Notas:           []float64{},
	}
}

// *Estudiante equivale a pasarle el puntero a la variable, y se usa para modificar los datos. En este caso modificamos el vector de notas
func (e *Estudiante) AgregarNota(nota float64) error {
	if nota < 0 || nota > 10 {
		return errors.New("nota inválida")
	}
	e.Notas = append(e.Notas, nota)
	return nil //equivalente a null
}

func (e *Estudiante) AgregarNotas(notas []float64) error {
	for _, n := range notas { //el for funciona con indice y valor, el _ indica que no nos interesa el primer valor, en este caso el indice. := hace que asigne de forma dinámica el tipo de dato
		if n < 0 || n > 10 {
			return errors.New("nota inválida en lista")
		}
		e.Notas = append(e.Notas, n)
	}
	return nil
}

// Como acá no necesitamos modificar la variable, no lleva *
func (e Estudiante) CalcularPromedio() float64 {
	if len(e.Notas) == 0 {
		return 0
	}
	acumulado := 0.0
	for _, nota := range e.Notas {
		acumulado += nota
	}
	return acumulado / float64(len(e.Notas))
}

// Método para verificar si es mayor de edad
func (e Estudiante) EsMayorDeEdad() bool {
	edad := time.Since(e.FechaNacimiento).Hours() / 24 / 365.25 //No hay algo como day o year, por eso usamos hour
	return edad >= 18
}
