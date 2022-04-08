package lsp

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
	Rectangle
}

func NewSquare(w float32) *Square {
	return &Square{Rectangle: Rectangle{w: w, h: w}}
}

func (s *Square) SetWidth(w float32) {
	s.w = w
	s.h = w
}

func (s *Square) SetHeight(h float32) {
	s.h = h
	s.w = h
}
