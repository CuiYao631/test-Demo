package main

import "fmt"

//  访问者模式

type Shape interface {
	accept(visitor)
}

type square struct{}

func (s *square) accept(v visitor) {
	v.visitForSquare(s)
}

type circle struct{}

func (c *circle) accept(v visitor) {
	v.visitForCircle(c)
}

type visitor interface {
	visitForSquare(*square)
	visitForCircle(*circle)
}

type sideCalculator struct{}

func (a *sideCalculator) visitForSquare(s *square) {
	fmt.Println("square side")
}

func (a *sideCalculator) visitForCircle(s *circle) {
	fmt.Println("circle side")
}

type radiusCalculator struct{}

func (a *radiusCalculator) visitForSquare(s *square) {
	fmt.Println("square radius")
}

func (a *radiusCalculator) visitForCircle(c *circle) {
	fmt.Println("circle radius")
}

func main() {
	square := &square{}
	circle := &circle{}

	side := &sideCalculator{}

	square.accept(side)
	circle.accept(side)

	radius := &radiusCalculator{}
	square.accept(radius)
	circle.accept(radius)

}
