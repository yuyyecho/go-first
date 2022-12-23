package main

import "fmt"

type Circle struct {
	shape Shape
	r     int
}

type Shape struct {
	draw Drawer
}
type Drawer interface {
	Draw(target string, x int, y int)
}

func (s *Shape) DrawShape(d Drawer) {
	s.draw = d
}

func (c *Circle) DrawCircle(r int, d Drawer) {
	c.shape.DrawShape(d)
	c.r = r
}
func (c *Circle) Draw(target string, r int, y int) {
	c.shape.draw.Draw("circle", c.r, 0)
}

type RedCircle struct {
}

func (r *RedCircle) Draw(target string, x int, y int) {
	fmt.Println("绘制一个红色的", target, "半径", x)
}

func main() {
	red := Circle{}
	red.DrawCircle(10, &RedCircle{})
	red.Draw("circle", 10, 0)
}
