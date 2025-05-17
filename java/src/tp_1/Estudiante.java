package tp_1;

import java.time.LocalDate;
import java.time.Period;
import java.util.ArrayList;
import java.util.List;
import java.util.Objects;

public class Estudiante implements Comparable<Estudiante> {
	private int dni;
	private String nombre;
	LocalDate fechaNacimiento;
	private List<Double> notas;

	public Estudiante(int dni, String nombre, LocalDate fechaNacimiento) {
		this.dni = dni;
		this.nombre = nombre;
		this.fechaNacimiento = fechaNacimiento;
		this.notas = new ArrayList<Double>();
	}

	public int getDni() {
		return this.dni;
	}

	public void agregarNota(double nota) {
		if (nota < 0 || nota > 10) {
			throw new IllegalArgumentException("NOTA INVALIDA");
		}
		this.notas.add(nota);
	}

	public void agregarNota(double[] notas) {
		for (int i = 0; i < notas.length; i++) {
			if (notas[i] < 0 || notas[i] > 10) {
				throw new IllegalArgumentException("NOTA INVALIDA");
			}
			this.notas.add(notas[i]);
		}
	}

	public double calcularPromedio() {
		double acumulado = 0;
		for (Double nota : this.notas) {
			acumulado += nota;
		}
		return acumulado / notas.size();
	}

	public boolean esMayorDeEdad() {
		int edad = Period.between(this.fechaNacimiento, LocalDate.now()).getYears();
		return edad >= 18;
	}

	@Override
	public int hashCode() {
		return Objects.hash(dni, fechaNacimiento, nombre, notas);
	}

	@Override
	public boolean equals(Object obj) {
		if (this == obj)
			return true;
		if (obj == null)
			return false;
		if (getClass() != obj.getClass())
			return false;
		Estudiante other = (Estudiante) obj;
		return dni == other.dni && Objects.equals(fechaNacimiento, other.fechaNacimiento)
				&& Objects.equals(nombre, other.nombre) && Objects.equals(notas, other.notas);
	}

	@Override
	public int compareTo(Estudiante o) {
		double diferenciaProm = this.calcularPromedio() - o.calcularPromedio();
		if (diferenciaProm > 0) {
			return -1;
		}
		if (diferenciaProm < 0) {
			return 1;
		}
		return this.dni - o.dni;
	}

	@Override
	public String toString() {
		return "Estudiante [dni=" + dni + ", nombre=" + nombre + ", fechaNacimiento=" + fechaNacimiento + ", notas="
				+ notas + "]";
	}
	
	
}
