package tests

import (
	"testing"
	"time"

	"gestorEstudiantesEnGo/estudiante"

	"github.com/stretchr/testify/assert"
)

func TestAgregarNotaNegativaFalla(t *testing.T) {
	e := estudiante.NuevoEstudiante(1, "Agustin", time.Date(2001, 5, 20, 0, 0, 0, 0, time.UTC))

	assert.EqualError(t, e.AgregarNota(-5), "nota inválida")
}

func TestAgregarNotaMayorA10Falla(t *testing.T) {
	e := estudiante.NuevoEstudiante(1, "Agustin", time.Date(2001, 5, 20, 0, 0, 0, 0, time.UTC))

	assert.EqualError(t, e.AgregarNota(20), "nota inválida")
}

func TestCalcularPromedioCalculaEsperado(t *testing.T) {
	e := estudiante.NuevoEstudiante(1, "Agustin", time.Date(2001, 5, 20, 0, 0, 0, 0, time.UTC))

	e.AgregarNotas([]float64{10, 9, 10, 9})

	assert.Equal(t, 9.5, e.CalcularPromedio())
}

func TestCalcularPromedioEstudianteSinNotasDevuelveCero(t *testing.T) {
	e := estudiante.NuevoEstudiante(1, "Agustin", time.Date(2001, 5, 20, 0, 0, 0, 0, time.UTC))

	assert.Equal(t, 0.0, e.CalcularPromedio())
}

func TestEstudianteDe24AniosEsMayorDeEdad(t *testing.T) {
	fecha := time.Now().AddDate(-24, 0, 0) // 24 años atrás desde hoy
	e := estudiante.NuevoEstudiante(1, "Agustin", fecha)

	assert.True(t, e.EsMayorDeEdad())
}

func TestEstudianteDe17AniosNoEsMayorDeEdad(t *testing.T) {
	fecha := time.Now().AddDate(-17, 0, 0) // 17 años atrás desde hoy
	e := estudiante.NuevoEstudiante(1, "Bart", fecha)

	assert.False(t, e.EsMayorDeEdad())
}

func TestEstudiantesConMismosAtributosSonIguales(t *testing.T) {
	e1 := estudiante.NuevoEstudiante(1, "Bart", time.Date(2008, 1, 1, 0, 0, 0, 0, time.UTC))
	e1.AgregarNotas([]float64{8.0})

	e2 := estudiante.NuevoEstudiante(1, "Bart", time.Date(2008, 1, 1, 0, 0, 0, 0, time.UTC))
	e2.AgregarNotas([]float64{8.0})

	assert.Equal(t, e1, e2)
}

func TestEstudiantesDistintosSonDistintos(t *testing.T) {
	e1 := estudiante.NuevoEstudiante(1, "Bart", time.Date(2008, 1, 1, 0, 0, 0, 0, time.UTC))
	e1.AgregarNotas([]float64{8.0})

	e2 := estudiante.NuevoEstudiante(1, "Homero", time.Date(1975, 1, 1, 0, 0, 0, 0, time.UTC))
	e2.AgregarNotas([]float64{7.0})

	assert.NotEqual(t, e1, e2)
}

func TestReferenciasAlMismoEstudianteSonElMismoObjeto(t *testing.T) {
	e1 := estudiante.NuevoEstudiante(1, "Agustin", time.Date(2001, 5, 20, 0, 0, 0, 0, time.UTC))
	otraReferencia := e1

	assert.Same(t, e1, otraReferencia)
}

func TestReferenciasAEstudiantesConMismosAtributosNoSonElMismoObjeto(t *testing.T) {
	e1 := estudiante.NuevoEstudiante(1, "Agustin", time.Date(2001, 5, 20, 0, 0, 0, 0, time.UTC))
	e2 := estudiante.NuevoEstudiante(1, "Agustin", time.Date(2001, 5, 20, 0, 0, 0, 0, time.UTC))

	assert.NotSame(t, e1, e2)
}
