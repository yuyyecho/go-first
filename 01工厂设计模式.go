package main

type Pener interface {
	Writer()
}

//type Pencil struct {
//}
//type BrushPen struct {
//}
//
//func (p *Pencil) Writer() {
//	fmt.Println("铅笔写字....")
//}
//
//func (b *BrushPen) Writer() {
//	fmt.Println("毛笔写字")
//}

type PenFactory struct {
}

func (f *PenFactory) Product(pen string) Pener {
	switch pen {
	case "pencil":
		return f.ProductPencil()
	case "brushpen":
		return f.ProductBrushPen()
	default:
		return nil
	}

}
func (f *PenFactory) ProductPencil() Pener {
	return new(Pencil)
}
func (f *PenFactory) ProductBrushPen() Pener {
	return new(BrushPen)
}

func main() {
	f := PenFactory{}
	Pencil := f.Product("pencil")
	Pencil.Writer()
}
