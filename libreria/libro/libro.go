package libro

import "fmt"

//definición de objeto de tipo libro, con sus atributos
type Libro struct {
	name    string
	autor   string
	paginas int
}

type LibroTexto struct {
	Libro
	editorial string
	nivel     string
}

//Similitud con un constructor, esta función devuelve un tipo de dato Libro
func NuevoLibro(name, autor string, paginas int) *Libro {
	return &Libro{
		name:    name,
		autor:   autor,
		paginas: paginas,
	}
}

func NuevoLibroTexto(name, autor string, paginas int, editorial, nivel string) *LibroTexto {
	return &LibroTexto{
		Libro:     Libro{name: name, autor: autor, paginas: paginas},
		editorial: editorial,
		nivel:     nivel,
	}
}
func (l *Libro) SetName(name string) {
	l.name = name
}

func (l *Libro) GetName() string {
	return l.name
}

// polimorfismo
//el mismo metodo para diferentes tipos de datos

//métodos del objeto de tipo Libro, en mayus para que sean publicos, y se puedan acceder desde afuera del package.
func (l *Libro) Imprimir_info() {
	fmt.Printf("Título: %s, autor: %s, Páginas: %d\n", l.name, l.autor, l.paginas)
}

func (l *LibroTexto) Imprimir_info() {
	fmt.Printf("Título: %s, autor: %s, Páginas: %d, Editorial: %s, Nivel: %s. \n", l.name, l.autor, l.paginas, l.editorial, l.nivel)
}
