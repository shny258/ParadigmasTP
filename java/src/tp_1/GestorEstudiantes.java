package tp_1;

import java.util.ArrayList;
import java.util.List;
import java.util.Set;
import java.util.TreeSet;

public class GestorEstudiantes {
	private Set<Estudiante> estudiantes;
	
	public GestorEstudiantes() {
		this.estudiantes = new TreeSet<>();
	}
	
	public GestorEstudiantes(Estudiante estudiante) {
		this();
		this.estudiantes.add(estudiante);
	}
	
	public GestorEstudiantes(Estudiante[] estudiantes) {
		this();
		for (int i = 0; i < estudiantes.length; i++) {
			this.estudiantes.add(estudiantes[i]);			
		}
	}
	
	public GestorEstudiantes(Set<Estudiante> estudiantes) {
		this();
		this.estudiantes.addAll(estudiantes);
	}
	
	public GestorEstudiantes(List<Estudiante> estudiantes) {
		this();
		this.estudiantes.addAll(estudiantes);
	}
	
	public Estudiante buscarEstudiantePorDNI(int dni) {
		for (Estudiante estudiante : this.estudiantes) {
			if(estudiante.getDni() == dni) {
				return estudiante;
			}
		}
		return null;
	}
	
	public List<Estudiante> listarEstudiantesPromedioAprobado(){
		List<Estudiante> estudiantesRet = new ArrayList<>();
		
		for (Estudiante estudiante : this.estudiantes) {
			if(estudiante.calcularPromedio() >= 7) {
				estudiantesRet.add(estudiante);
			}
		}
		return estudiantesRet;
	}
}

