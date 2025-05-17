package tp_1;

import static org.junit.jupiter.api.Assertions.*;
import java.time.LocalDate;
import java.util.ArrayList;
import java.util.List;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;

class GestorEstudiantesTests {
	GestorEstudiantes gestor;
	List<Estudiante> estudiantes;

	@BeforeEach
	void setUp() {
		estudiantes = new ArrayList<>();

		Estudiante e0 = new Estudiante(1000, "Ana López", LocalDate.of(2003, 5, 15));
		e0.agregarNota(new double[] { 8.5, 7.0, 9.0, 6.5, 8.0 });

		Estudiante e1 = new Estudiante(1001, "Bruno Díaz", LocalDate.of(2000, 10, 10));
		e1.agregarNota(new double[] { 6.0, 7.5, 7.0, 8.0, 6.5 });

		Estudiante e2 = new Estudiante(1002, "Carla Gómez", LocalDate.of(2002, 3, 22));
		e2.agregarNota(new double[] { 9.0, 8.5, 9.5, 8.0, 9.0 });

		Estudiante e3 = new Estudiante(1003, "Diego Pérez", LocalDate.of(2004, 12, 3));
		e3.agregarNota(new double[] { 5.0, 6.0, 5.5, 6.5, 7.0 });

		Estudiante e4 = new Estudiante(1004, "Elena Torres", LocalDate.of(2001, 7, 28));
		e4.agregarNota(new double[] { 7.0, 7.5, 8.0, 8.5, 7.5 });

		Estudiante e5 = new Estudiante(1005, "Federico Ruiz", LocalDate.of(2003, 9, 9));
		e5.agregarNota(new double[] { 6.0, 6.5, 6.5, 7.0, 6.0 });

		Estudiante e6 = new Estudiante(1006, "Gabriela Castro", LocalDate.of(2000, 1, 30));
		e6.agregarNota(new double[] { 9.5, 9.0, 9.5, 10.0, 9.0 });

		Estudiante e7 = new Estudiante(1007, "Hernán Vidal", LocalDate.of(2002, 6, 6));
		e7.agregarNota(new double[] { 4.0, 5.5, 6.0, 5.0, 5.0 });

		Estudiante e8 = new Estudiante(1008, "Isabel Méndez", LocalDate.of(2001, 11, 19));
		e8.agregarNota(new double[] { 8.0, 8.0, 8.5, 8.5, 9.0 });

		Estudiante e9 = new Estudiante(1009, "Joaquín Herrera", LocalDate.of(2004, 2, 14));
		e9.agregarNota(new double[] { 7.5, 6.5, 7.0, 7.0, 6.0 });

		estudiantes.add(e0);
		estudiantes.add(e1);
		estudiantes.add(e2);
		estudiantes.add(e3);
		estudiantes.add(e4);
		estudiantes.add(e5);
		estudiantes.add(e6);
		estudiantes.add(e7);
		estudiantes.add(e8);
		estudiantes.add(e9);

		gestor = new GestorEstudiantes(estudiantes);
	}

	@Test
	void buscarEstudianteExistenteLoEncuentra() {
		assertSame(this.estudiantes.get(7), gestor.buscarEstudiantePorDNI(1007));
	}

	@Test
	void buscarEstudianteInexistenteNoLoEncuentra() {
		assertNull(gestor.buscarEstudiantePorDNI(1010));
	}

	@Test
	void listarEstudiantesPromedioAprobadoListaAprobados() {
		Estudiante[] aprobadosEsperados = { estudiantes.get(6), estudiantes.get(2), estudiantes.get(8),
				estudiantes.get(0), estudiantes.get(4), estudiantes.get(1) };

		assertArrayEquals(aprobadosEsperados, gestor.listarEstudiantesPromedioAprobado().toArray());
	}

}
