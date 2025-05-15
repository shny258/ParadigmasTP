package tp_1;

import static org.junit.Assert.assertThrows;

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

}
