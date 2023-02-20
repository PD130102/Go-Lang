package main

import ("fmt"
"math")

type Shape interface {
	area() float64
	perimeter() float64
}

type Square struct {
	side float64
}

type Triangle struct {
   side1, side2, side3 float64
}

type Circle struct {
   radius float64
}

type Rectangle struct {
   width, height float64
}

func (square Square) area() float64 {
	return square.side * square.side
}

func (circle Circle) area() float64{
	return math.Pi * circle.radius * circle.radius
}

func (triangle Triangle) area() float64 {
	return 0.5 * triangle.side1 * triangle.side2
}

func (rectangle Rectangle) area() float64 {
	return rectangle.width * rectangle.height
}

func (rectangle Rectangle) perimeter() float64 {
	return 2 * (rectangle.width + rectangle.height)
}

func (triangle Triangle) perimeter() float64{
	return triangle.side1 + triangle.side2 + triangle.side3
}

func (square Square) perimeter() float64 {
	return 4 * square.side
}

func (circle Circle) perimeter() float64 {
	return 2 * math.Pi * circle.radius
}

func getArea(shape Shape) float64 {
	return shape.area()
}
func getPerimeter(shape Shape) float64 {
	return shape.perimeter()
}

func main(){
	square := Square{side:5}
	circle := Circle{radius:5}
	triangle := Triangle{side1:5, side2:5, side3:5}
	rectangle := Rectangle{width:5, height:5}

	fmt.Printf("Square area: %f\n", square.area())
	fmt.Printf("Circle area: %f\n", circle.area())
	fmt.Printf("Triangle area: %f\n", triangle.area())
	fmt.Printf("Rectangle area: %f\n", rectangle.area())
	fmt.Printf("Rectangle perimeter: %f\n", rectangle.perimeter())
	fmt.Printf("Triangle perimeter: %f\n", triangle.perimeter())
	fmt.Printf("Square perimeter: %f\n", square.perimeter())
	fmt.Printf("Circle perimeter: %f\n", circle.perimeter())

}