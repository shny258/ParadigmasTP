package tests

import (
	"testing"
	"time"

	"gestorEstudiantesEnGo/estudiante"

	"github.com/stretchr/testify/assert"
)

func crearEstudiantes() ([]*estudiante.Estudiante, *estudiante.GestorEstudiantes) {
	estudiantes := []*estudiante.Estudiante{
		estudiante.NuevoEstudiante(1000, "Ana López", time.Date(2003, 5, 15, 0, 0, 0, 0, time.UTC)),
		estudiante.NuevoEstudiante(1001, "Bruno Díaz", time.Date(2000, 10, 10, 0, 0, 0, 0, time.UTC)),
		estudiante.NuevoEstudiante(1002, "Carla Gómez", time.Date(2002, 3, 22, 0, 0, 0, 0, time.UTC)),
		estudiante.NuevoEstudiante(1003, "Diego Pérez", time.Date(2004, 12, 3, 0, 0, 0, 0, time.UTC)),
		estudiante.NuevoEstudiante(1004, "Elena Torres", time.Date(2001, 7, 28, 0, 0, 0, 0, time.UTC)),
		estudiante.NuevoEstudiante(1005, "Federico Ruiz", time.Date(2003, 9, 9, 0, 0, 0, 0, time.UTC)),
		estudiante.NuevoEstudiante(1006, "Gabriela Castro", time.Date(2000, 1, 30, 0, 0, 0, 0, time.UTC)),
		estudiante.NuevoEstudiante(1007, "Hernán Vidal", time.Date(2002, 6, 6, 0, 0, 0, 0, time.UTC)),
		estudiante.NuevoEstudiante(1008, "Isabel Méndez", time.Date(2001, 11, 19, 0, 0, 0, 0, time.UTC)),
		estudiante.NuevoEstudiante(1009, "Joaquín Herrera", time.Date(2004, 2, 14, 0, 0, 0, 0, time.UTC)),
	}

	notas := [][]float64{
		{8.5, 7.0, 9.0, 6.5, 8.0},
		{6.0, 7.5, 7.0, 8.0, 6.5},
		{9.0, 8.5, 9.5, 8.0, 9.0},
		{5.0, 6.0, 5.5, 6.5, 7.0},
		{7.0, 7.5, 8.0, 8.5, 7.5},
		{6.0, 6.5, 6.5, 7.0, 6.0},
		{9.5, 9.0, 9.5, 10.0, 9.0},
		{4.0, 5.5, 6.0, 5.0, 5.0},
		{8.0, 8.0, 8.5, 8.5, 9.0},
		{7.5, 6.5, 7.0, 7.0, 6.0},
	}

	for i, est := range estudiantes {
		est.AgregarNotas(notas[i])
	}

	gestor := estudiante.NuevoGestorConSlice(estudiantes)

	return estudiantes, gestor
}

func TestBuscarEstudianteExistenteLoEncuentra(t *testing.T) {
	estudiantes, gestor := crearEstudiantes()
	resultado := gestor.BuscarEstudiantePorDNI(1007)
	assert.Same(t, estudiantes[7], resultado)
}

func TestBuscarEstudianteInexistenteNoLoEncuentra(t *testing.T) {
	_, gestor := crearEstudiantes()
	resultado := gestor.BuscarEstudiantePorDNI(1010)
	assert.Nil(t, resultado)
}

func TestListarEstudiantesPromedioAprobadoListaAprobados(t *testing.T) {
	estudiantes, gestor := crearEstudiantes()

	esperados := []*estudiante.Estudiante{
		estudiantes[6],
		estudiantes[2],
		estudiantes[8],
		estudiantes[0],
		estudiantes[4],
		estudiantes[1],
	}

	obtenidos := gestor.ListarEstudiantesPromedioAprobado()
	assert.ElementsMatch(t, esperados, obtenidos) //NO tiene en cuenta el orden
}
