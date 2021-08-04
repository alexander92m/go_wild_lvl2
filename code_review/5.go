// Что выведет программа? Объяснить вывод программы.
package main

import "fmt"
type customError struct {
	msg string
}

func (e *customError) Error() string {
	return e.msg
}
func test() *customError {
	{
	// do something
	}
	return nil
}

func main() {
	var err error
	//интерфейс содержащий нулевой указатель не равен nil
	err = test() 
	
	//а вот если возвращать структуру то всё ок
	// err2 := test()
	//
	// fmt.Printf("%#v\n", err)
	// fmt.Printf("%#v\n", err2)
	// if err2 != nil {
	// 	println("error")
	// 	return
	// }
	if err != nil {
		println("error")
		return
	}
	println("ok")
}