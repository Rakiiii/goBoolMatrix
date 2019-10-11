package boolmatrixlib

import "math"

//BoolMatrix is struct that simulate matrix of bools in linair massive with per bit value storage
type BoolMatrix struct {
	matrix       []byte
	width, heigh int
}

func (m *BoolMatrix) Init(w, h int) {
	m.matrix = make([]byte, int64(math.Ceil((float64(w)*float64(h))/8)))
	m.width = w
	m.heigh = h
}
