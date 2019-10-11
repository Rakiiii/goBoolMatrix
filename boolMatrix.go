package boolmatrixlib

import "math"

//BoolMatrix is struct that simulate matrix of bools in linair massive with per bit value storage
type BoolMatrix struct {
	matrix       []byte
	width, heigh int
}

//Init is function of initialization for BoolMatrix
func (m *BoolMatrix) Init(w, h int) {
	m.matrix = make([]byte, int64(math.Ceil((float64(w)*float64(h))/8)))
	m.width = w
	m.heigh = h
}

//GetBool is immulating index operator for matrix returning value containing in @i row and j columne
func (m *BoolMatrix) GetBool(i, j int) bool {
	if i > m.heigh || j > m.width {
		return false
	}
	var it = m.matrix[(i*m.width+j)/8]
	switch (i*m.width + j) % 8 {
	case 0:
		return (it & 128) > 0
	case 1:
		return (it & 64) > 0
	case 2:
		return (it & 32) > 0
	case 3:
		return (it & 16) > 0
	case 4:
		return (it & 8) > 0
	case 5:
		return (it & 4) > 0
	case 6:
		return (it & 2) > 0
	case 7:
		return (it & 1) > 0
	}
	return false
}

func (m *BoolMatrix) SetBool(i, j int, value bool) bool {
	if i > m.heigh || j > m.width {
		return false
	}
	if value {
		m.matrix[(i*m.width+j)/8] = (m.matrix[(i*m.width+j)/8] & 255)
	} else {
		switch (i*m.width + j) % 8 {
		case 0:
			m.matrix[(i*m.width+j)/8] = (m.matrix[(i*m.width+j)/8] & 127)
			return true
		case 1:
			m.matrix[(i*m.width+j)/8] = (m.matrix[(i*m.width+j)/8] & 191)
			return true
		case 2:
			m.matrix[(i*m.width+j)/8] = (m.matrix[(i*m.width+j)/8] & 223)
			return true
		case 3:
			m.matrix[(i*m.width+j)/8] = (m.matrix[(i*m.width+j)/8] & 239)
			return true
		case 4:
			m.matrix[(i*m.width+j)/8] = (m.matrix[(i*m.width+j)/8] & 247)
			return true
		case 5:
			m.matrix[(i*m.width+j)/8] = (m.matrix[(i*m.width+j)/8] & 251)
			return true
		case 6:
			m.matrix[(i*m.width+j)/8] = (m.matrix[(i*m.width+j)/8] & 253)
			return true
		case 7:
			m.matrix[(i*m.width+j)/8] = (m.matrix[(i*m.width+j)/8] & 254)
			return true
		}
	}
	return false
}

func (m *BoolMatrix) SetByNumber(number int) bool {
	if math.Pow(2, float64(m.width*m.heigh))-1 < float64(number) {
		return false
	}

}
