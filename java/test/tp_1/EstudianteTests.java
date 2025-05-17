package tp_1;

import static org.junit.jupiter.api.Assertions.*;
import java.time.LocalDate;
import org.junit.jupiter.api.Test;

class EstudianteTests {
	Estudiante estudiante;
	
	@Test
	void agregarNotaNegativaFalla() {
		estudiante = new Estudiante(1, "Agustin", LocalDate.of(2001, 5, 20));
		assertThrows(IllegalArgumentException.class, () -> {
			estudiante.agregarNota(-5);
		});
	}
	
	@Test
	void agregarNotaMayorA10Falla() {
		estudiante = new Estudiante(1, "Agustin", LocalDate.of(2001, 5, 20));
		assertThrows(IllegalArgumentException.class, () -> {
			estudiante.agregarNota(20);
		});
	}
	
	@Test
	void calcularPromedioCalculaEsperado() {
		estudiante = new Estudiante(1, "Agustin", LocalDate.of(2001, 5, 20));
		double[] notas = {10,9,10,9};
		estudiante.agregarNota(notas);
		assertEquals(9.5, estudiante.calcularPromedio());
	}
	
	@Test
	void calcularPromedioEstudianteSinNotasFalla() {
		estudiante = new Estudiante(1, "Agustin", LocalDate.of(2001, 5, 20));
		assertTrue(Double.isNaN(estudiante.calcularPromedio()));
	}
	
	@Test
	void estudianteDe24AñosEsMayorDeEdad() {
		estudiante = new Estudiante(1, "Agustin", LocalDate.of(2001, 5, 20));
		assertTrue(estudiante.esMayorDeEdad());
	}
	
	@Test
	void estudianteDe17AñosNoEsMayorDeEdad() {
		estudiante = new Estudiante(1, "Bart", LocalDate.of(2008, 1, 1));
		assertFalse(estudiante.esMayorDeEdad());
	}
	
	@Test
	void estudiantesConMismosAtributosSonIguales() {
		estudiante = new Estudiante(1, "Bart", LocalDate.of(2008, 1, 1));
		Estudiante estudiante2 = new Estudiante(1, "Bart", LocalDate.of(2008, 1, 1));
		assertEquals(estudiante, estudiante2);
	}
	
	@Test
	void estudiantesDistintosSonDistintos() {
		estudiante = new Estudiante(1, "Bart", LocalDate.of(2008, 1, 1));
		Estudiante estudiante2 = new Estudiante(1, "Homero", LocalDate.of(1975, 1, 1));
		assertNotEquals(estudiante, estudiante2);
	}
	
	@Test
	void referenciasAlMismoEstudianteSonElMismoObjeto(){
		estudiante = new Estudiante(1, "Agustin", LocalDate.of(2001, 5, 20));
		Estudiante otraReferencia = estudiante;
		assertSame(estudiante,otraReferencia);
	}
	
	@Test
	void referenciasAEstudiantesConMismosAtributosNoSonElMismoObjeto(){
		estudiante = new Estudiante(1, "Agustin", LocalDate.of(2001, 5, 20));
		Estudiante otraReferencia = new Estudiante(1, "Agustin", LocalDate.of(2001, 5, 20));
		assertNotSame(estudiante,otraReferencia);
	}

}
