package Funciones

import "fmt"

func MostrarNumeros() {
	/**
	 * Formas para declarar variables.
	 * 1. Var -> Al proporcionar la palabra clave var, le estamos indicando que va a tener un nombre propio
	 * 	{var Identificador/nombre tipo = (valor inicial) - opcional}
	 */
	var intcomun int

	/**
	 * Declaracion corta de variables.
	 * Consiste en asignar el tipo de variable con base al contenido, si el valor es un numero, la variable sera
	 * de tipo int, si es una cadena sera un string. Generalmente se usa dentro de las funciones.
	 *
	 * En este caso por ejemplo las funciones int convierte un valor en su equivalencia, en este caso
	 * 32 y 64. Por lo que valor devuelto sera un numero, y en consecuencia las variables seran de tipo int.
	 */
	entero32 := int32(10)
	entero64 := int64(100)

	fmt.Println("Salida para intComun: ", intcomun)
	fmt.Println("Salida para entero32: ", entero32)
	fmt.Println("Salida para entero64: ", entero64)
}
