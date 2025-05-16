package Estudiante

// GestorEstudiantes usa un map con clave DNI
type GestorEstudiantes struct {
	estudiantes map[int]*Estudiante
}

// Constructor vacío
func NuevoGestorEstudiantes() *GestorEstudiantes {
	return &GestorEstudiantes{
		estudiantes: make(map[int]*Estudiante),
	}
}

// En go no existe la sobrecarga, asi que debemos crear dos métodos distintos
func NuevoGestorConEstudiante(e *Estudiante) *GestorEstudiantes {
	g := NuevoGestorEstudiantes()
	g.estudiantes[e.DNI] = e
	return g
}

// Constructor con slice de estudiantes
func NuevoGestorConSlice(slice []*Estudiante) *GestorEstudiantes {
	g := NuevoGestorEstudiantes()
	for _, e := range slice {
		g.estudiantes[e.DNI] = e
	}
	return g
}

// Buscar estudiante por DNI
func (g *GestorEstudiantes) BuscarEstudiantePorDNI(dni int) *Estudiante {
	if est, existe := g.estudiantes[dni]; existe {
		return est
	}
	return nil
}

// Listar estudiantes con promedio >= 7
func (g *GestorEstudiantes) ListarEstudiantesPromedioAprobado() []*Estudiante {
	var aprobados []*Estudiante
	for _, e := range g.estudiantes {
		if e.CalcularPromedio() >= 7 {
			aprobados = append(aprobados, e)
		}
	}
	return aprobados
}
