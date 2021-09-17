// Что выведет программа? Объяснить вывод программы. Объяснить внутреннее
// устройство интерфейсов и их отличие от пустых интерфейсов.

package main

import (
	"fmt"
	"os"
)

func Foo() error {
	var err *os.PathError = nil
	return err
}

func main() {
	err := Foo()

	fmt.Println(err)
	fmt.Println(err == nil)
	// переменная равна nil только если и тип переменной и значение равны nil
	fmt.Printf("1 %#v, %#v\n", err, nil)


	// Интерфейс-это набор методов, а не набор полей. Тип реализует интерфейс, 
	//если его методы включают методы этого интерфейса. Поскольку пустой 
	//интерфейс не имеет никаких методов, все типы реализуют его.
	var v interface{}
	fmt.Printf("2 %T %v %v\n", v, v, v == nil)
	//использовать пустой интерфейс удобно в случаях когда функция должна обработать любой тип переменной в рантайме
	var i int
	v = &i
	fmt.Printf("3 %T %v %v\n", v, v, v == nil)
	var a []int
	v = a
	fmt.Printf("4 %T %v %v\n", v, v, v == nil)
	fmt.Printf("5 %T %v %v\n", a, a, a == nil)
	//под капотом элемент интерфейса указывает на тип и значение, это позволяет
	//  ему быть многоликим, но оператор сравнения считает что интерфейс
	//   равен nil только если тип и значение элемента интерфейса nil
}