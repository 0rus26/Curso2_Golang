package main

import (
	"fmt"
	"time"
)

func main() {
	lista := listaTareas{}
	fmt.Println(">--INICIO--<")
	var salir int
	for {

		if salir != 1 {
			opcion := control()
			//fmt.Println("Opción seleccionada:", opcion)
			switch opcion {
			case 1:
				lista.case1()
			case 2:
				lista.case2()
			case 3:
				lista.case3()
			case 4:
				lista.case4()
			case 5:
				lista.case5()
			case 6:
				salir = 1
			default:
				fmt.Println("Opción no disponible")
				break
			}
		} else {
			fmt.Println("Terminando ejecución o7 ")
			break
		}

		//fmt.Println("fin-->")
	}

}

func control() int {
	var opcion int
	fmt.Println("\n Selecionar opción:\n",
		"1.add\n",
		"2.marcar done\n",
		"3.edit\n",
		"4.delete\n",
		"5.Ver tareas\n",
		"6.exit\n")
	fmt.Println("Ingrese la opción:")
	fmt.Scan(&opcion)
	fmt.Println("\n")
	return opcion
}

type Tarea struct {
	id          int
	nombre      string
	descripcion string
	fecha       string
	completado  bool
}

type listaTareas struct {
	tareas []Tarea
}

func (l *listaTareas) add_Tareas(t Tarea) {

	now := time.Now()
	nueva_Tarea := Tarea{
		id:          len(l.tareas) + 1,
		nombre:      t.nombre,
		descripcion: t.descripcion,
		fecha:       now.String(),
		completado:  false,
	}
	l.tareas = append(l.tareas, nueva_Tarea)

}

func (l *listaTareas) delete_Tareas(i int) {
	l.tareas = append(l.tareas[:i], l.tareas[i+1:]...)

}

func (l *listaTareas) update_Tareas(i int, t Tarea) {
	l.tareas[i] = t
}

func (l *listaTareas) done_Tareas(i int) {
	if len(l.tareas) == 0 {
		fmt.Println("No hay tareas")
	} else {
		l.readAll_Task()
		fmt.Println("Seleccione la tarea a finalizar:")
		var item int
		fmt.Scan(&item)
		fmt.Println("\n")
		item--
		l.tareas[i].completado = true
		fmt.Println("Tarea ", l.tareas[i].nombre, "ha sido finalizada")
	}
}

func (l *listaTareas) readAll_Task() {

	if len(l.tareas) == 0 {
		fmt.Println(">--NO HAY TAREAS DISPONIBLES--<")
	} else {
		fmt.Println(">--LISTA_DE_TAREAS--<")
		for _, tareas := range l.tareas {
			estado := ""
			switch tareas.completado {
			case true:
				estado = "DONE"
			case false:
				estado = "IN_PROGRESS"
			}
			fmt.Printf("TAREA#%d-> Nombre:%s Estado: %v\n", tareas.id, tareas.nombre, estado)
		}
	}

}

//CASES main

func (l *listaTareas) case1() {
	var t Tarea
	fmt.Printf("\nDigite nombre:\n")
	fmt.Scan(&t.nombre)
	fmt.Println("\n")
	fmt.Printf("\nDigite descripcion:\n")
	fmt.Scan(&t.descripcion)
	fmt.Println("\n")
	fmt.Println("Tarea", t.nombre, "fue agregada correctamente.")
	//fmt.Println("<<-->>")
	l.add_Tareas(t)
}

func (l *listaTareas) case2() {
	var item int
	l.readAll_Task()

	if len(l.tareas) == 0 {
		fmt.Println(">--ERROR: No es posible marcar tareas--<")
		return
	}
	fmt.Println("Seleccione el numero de la tarea a terminar (int): ")
	fmt.Scan(&item)
	fmt.Println("\n")

	item--
	l.done_Tareas(item)

	l.readAll_Task()
}

func (l *listaTareas) case3() {
	var t Tarea
	var item int

	fmt.Println("Seleccione la tarea que desea actualizar (int): ")
	fmt.Scan(&item)
	fmt.Println("\n")
	fmt.Println("Indique el nombre de la tarea: ")
	fmt.Scan(&t.nombre)
	fmt.Println("\n")
	fmt.Println("Indique la descripción de la tarea: ")
	fmt.Scan(&t.descripcion)
	fmt.Println("\n")
	item--
	l.update_Tareas(item, t)
	fmt.Println("Tarea actualizada correctamente: ")
	fmt.Println("\n")
}

func (l *listaTareas) case4() {
	var item int
	fmt.Println("Seleccione la tarea que desea eliminar (int): ")
	fmt.Scan(&item)
	item--
	l.delete_Tareas(item)
	fmt.Println("Tarea eliminada correctamente: ")
	fmt.Println("\n")
}

func (l *listaTareas) case5() {
	l.readAll_Task()

}
