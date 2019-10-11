package boolmatrixlib

import (
	"math"
	"math/big"
)

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

//Width return matrix width
func (m *BoolMatrix) Width() int {
	return m.width
}

//Heigh return matrix heigh
func (m *BoolMatrix) Heigh() int {
	return m.heigh
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

func (m *BoolMatrix) SetByNumber(num *big.Int) bool {
	slice := num.Bytes()
	l := len(slice)
	if l > len(m.matrix) {
		return false
	}
	dif := len(m.matrix) - l
	for i := 0; i < l; i++ {
		m.matrix[i+dif] = slice[i]
	}
	if dif != 0 {
		for i := 0; i < dif; i++ {
			m.matrix[i] = 0
		}
	}
	return true
}

func (m *BoolMatrix) CountTrues() int64 {
	counter := int64(0)
	for _, i := range m.matrix {
		counter += int64((i >> 7) + ((i >> 6) & 1) + ((i >> 5) & 1) + ((i >> 4) & 1) + ((i >> 3) & 1) + ((i >> 2) & 1) + ((i >> 1) & 1))
	}
	return counter
}

func (m *BoolMatrix) CountTruesInLine(line int) int {
	lineStart := line * m.width / 8
	counter := line * m.width % 8
	res := 0
	for i := 0; i < m.width; i++ {
		res += int(m.matrix[lineStart+counter/8] & byte(math.Pow(2, float64(counter%8))))
		counter++
	}
	return res
}

func (m *BoolMatrix) Copy() *BoolMatrix {
	newMatrix := new(BoolMatrix)
	newMatrix.Init(m.width, m.heigh)
	for i := 0; i < len(m.matrix); i++ {
		newMatrix.matrix[i] = m.matrix[i]
	}
	return newMatrix
}
