package main

import "fmt"


type visitor interface {
    visitForSquare(*square)
    visitForCircle(*circle)
    visitForRectangle(*rectangle)
}

//square квадрат
type square struct {
    side int
}

//rectangle треугольник
type rectangle struct {
    l int
    b int
}

//circle круг
type circle struct {
    radius int
}

//shape методы для фигур
type shape interface {
    getType() string
    accept(visitor)
}

//middleCoordinates координаты центра
type middleCoordinates struct {
    x int
    y int
}

//areaCalculator площадь фигуры
type areaCalculator struct {
    area int
}

func main() {
    square := &square{side: 2}
    circle := &circle{radius: 3}
    rectangle := &rectangle{l: 2, b: 3}

    areaCalculator := &areaCalculator{}

    square.accept(areaCalculator)
    circle.accept(areaCalculator)
    rectangle.accept(areaCalculator)

    fmt.Println()
    middleCoordinates := &middleCoordinates{}
    square.accept(middleCoordinates)
    circle.accept(middleCoordinates)
    rectangle.accept(middleCoordinates)
}

//accept выбор фигуры
func (s *square) accept(v visitor) {
    v.visitForSquare(s)
}

//accept выбор фигуры
func (t *rectangle) accept(v visitor) {
    v.visitForRectangle(t)
}

//accept выбор фигуры
func (c *circle) accept(v visitor) {
    v.visitForCircle(c)
}

//getType определение типа
func (s *square) getType() string {
    return "Square"
}

//getType определение типа
func (t *rectangle) getType() string {
    return "rectangle"
}

//getType определение типа
func (c *circle) getType() string {
    return "Circle"
}

//visitForSquare вычисление координат центра квадрата
func (a *middleCoordinates) visitForSquare(s *square) {
    // Calculate middle point coordinates for square.
    // Then assign in to the x and y instance variable.
    fmt.Println("Calculating middle point coordinates for square")
}

//visitForCircle вычисление координат центра круга
func (a *middleCoordinates) visitForCircle(c *circle) {
    fmt.Println("Calculating middle point coordinates for circle")
}

//visitForRectangle вычисление координат центра треугольника
func (a *middleCoordinates) visitForRectangle(t *rectangle) {
    fmt.Println("Calculating middle point coordinates for rectangle")
}

//visitForSquares случай для квадрата
func (a *areaCalculator) visitForSquare(s *square) {
    // Calculate area for square.
    // Then assign in to the area instance variable.
    fmt.Println("Calculating area for square")
}

//visitForCircle случай для круга
func (a *areaCalculator) visitForCircle(s *circle) {
    fmt.Println("Calculating area for circle")
}

//visitForRectangle случай для треугольника
func (a *areaCalculator) visitForRectangle(s *rectangle) {
    fmt.Println("Calculating area for rectangle")
}