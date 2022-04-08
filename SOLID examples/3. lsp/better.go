package main

type RectangularShape interface {
	Area() float32

	Width() float32
	Height() float32

	SetWidth(w float32)
	SetHeight(w float32)
}

type Rectangle struct {
	w float32
	h float32
}

func (r *Rectangle) Area() float32 {
	return r.w * r.h
}

func (r *Rectangle) Width() float32 {
	return r.w
}

func (r *Rectangle) Height() float32 {
	return r.h
}

func (r *Rectangle) SetWidth(w float32) {
	r.w = w
}

func (r *Rectangle) SetHeight(h float32) {
	r.h = h
}

type Square struct {
	w float32
}

func (s *Square) Area() float32 {
	return s.w * s.w
}

func (s *Square) SetWidth(w float32) {
	s.w = w
}

type Folder struct{}

func (f Folder) FoldInHalf(r RectangularShape) RectangularShape {
	r.SetWidth(r.Width() / 2)
	return r
}

func main() {
	s := &Square{w: 2}

	f := Folder{}
	f.FoldInHalf(s) // doesn't compile (rightfully) as the Folder can't fold squares in half
}
