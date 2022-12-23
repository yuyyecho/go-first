package main

import "fmt"

type Pen interface {
	Writer()
}

// 抽象工厂
type AbstractFactory interface {
	Produce() Pen
}

type PencilFactory struct {
}

type Pencil struct {
}

func (p *Pencil) Writer() {
	fmt.Println("铅笔写字....")
}

func (p *PencilFactory) Product() Pen {
	return new(Pencil)
}

type BrushPenFactory struct {
}

type BrushPen struct {
}

func (b *BrushPen) Writer() {
	fmt.Println("毛笔写字")
}

func (p *BrushPenFactory) Product() Pen {
	return new(BrushPen)
}

func main() {
	f := PencilFactory{}
	Pencil := f.Product()
	Pencil.Writer()

	b := BrushPenFactory{};
	BrushPen := b.Product()
	BrushPen.Writer()
}
