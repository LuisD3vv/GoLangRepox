package main

import (
	"bufio"
	"fmt"
	"os"
)

type Tarea struct {
	nombre     string
	desc       string
	completado bool
}
type ListaTareas struct {
	tareas []Tarea
}

// Metodo para agregar tarea
func (l *ListaTareas) agregarTarea(t Tarea) {
	l.tareas = append(l.tareas, t)
}

// metodo para maracar tarea como completada
func (l *ListaTareas) marcarCompletado(index int) {
	l.tareas[index].completado = true
}

// metodo para editar
func (l *ListaTareas) editarTarea(index int, t Tarea) {
	l.tareas[index] = t
}

// metodo para eliminar tarea
func (l *ListaTareas) eliminarTarea(index int) {
	l.tareas = append(l.tareas[:index], l.tareas[index+1:]...)
}

func main() {
	// Instancia de Lista de tareas
	lista := ListaTareas{}
	//instancia de bufio para entrada de datos
	leer := bufio.NewReader(os.Stdin)
	for {
		var opcion int
		fmt.Println("Seleccione una opcion\n",
			"1. Agregar Tarea\n",
			"2. Marcar tarea como completada\n",
			"3. Editar Tarea\n",
			"4. Eliminar Tarea\n",
			"5. Salir")
		fmt.Println("Ingrese la opcion")
		fmt.Scanln(&opcion)
		switch opcion {
		case 1:
			var t Tarea
			fmt.Print("Ingrese el nombre de la tarea: ")
			t.nombre, _ = leer.ReadString('\n')
			fmt.Print("Ingrese la tarea: ")
			t.desc, _ = leer.ReadString('\n')
			lista.agregarTarea(t)
			fmt.Println("Tarea Agregada correctamente")

		case 2:
			var index int
			fmt.Print("Ingrese el indice de la tarea que desea marcar como completada: ")
			fmt.Scanln(&index)
			lista.marcarCompletado(index)
			fmt.Println("Tarea marcada como completada exitosamente")
		case 3:
			var index int
			var t Tarea
			fmt.Print("Ingrese el indice de la tarea para actualizar: ")
			fmt.Scanln(&index)
			fmt.Print("Ingrese el nombre de la tarea: ")
			t.nombre, _ = leer.ReadString('\n')
			fmt.Print("Ingrese la tarea: ")
			t.desc, _ = leer.ReadString('\n')
			lista.editarTarea(index, t)
			fmt.Println("Tarea Actualizada corectamente.")
		case 4:
			var index int
			fmt.Print("Ingrese el indice de la tarea que desea eliminar: ")
			fmt.Scanln(&index)
			lista.eliminarTarea(index)
		case 5:
			fmt.Println("Saliendo del programa...")
			return
		default:
			fmt.Println("Opcion invalida")
		}
		//Listar tareas
		fmt.Println("Lista de Tareas")
		fmt.Println("====================")
		for i, t := range lista.tareas {
			fmt.Printf("%d. %s - %s - Completado: %t\n", i, t.nombre, t.desc, t.completado)
		}
		fmt.Println("====================")
	}

}
