package boolmatrixlib

import (
	"fmt"
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

//SetMatrix is setting BoolMatrix equals to bool massive
func (m *BoolMatrix) SetMatrix(newmat []byte) {
	//	for i := 0; i < len(newmat); i++ {
	//		m.matrix[i] = newmat[i]
	//	}
	m.matrix = newmat
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
		//m.matrix[(i*m.width+j)/8] = (m.matrix[(i*m.width+j)/8] & 255)
		switch (i*m.width + j) % 8 {
		case 0:
			m.matrix[(i*m.width+j)/8] = (m.matrix[(i*m.width+j)/8] | 128)
			return true
		case 1:
			m.matrix[(i*m.width+j)/8] = (m.matrix[(i*m.width+j)/8] | 64)
			return true
		case 2:
			m.matrix[(i*m.width+j)/8] = (m.matrix[(i*m.width+j)/8] | 32)
			return true
		case 3:
			m.matrix[(i*m.width+j)/8] = (m.matrix[(i*m.width+j)/8] | 16)
			return true
		case 4:
			m.matrix[(i*m.width+j)/8] = (m.matrix[(i*m.width+j)/8] | 8)
			return true
		case 5:
			m.matrix[(i*m.width+j)/8] = (m.matrix[(i*m.width+j)/8] | 4)
			return true
		case 6:
			m.matrix[(i*m.width+j)/8] = (m.matrix[(i*m.width+j)/8] | 2)
			return true
		case 7:
			m.matrix[(i*m.width+j)/8] = (m.matrix[(i*m.width+j)/8] | 1)
			return true
		}
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
	if len(m.matrix)*8 > m.width*m.heigh {
		dif = 8 - (m.width*m.heigh)%8
		var bitMask byte
		switch dif {
		case 1:
			bitMask = 128
		case 2:
			bitMask = 192
		case 3:
			bitMask = 224
		case 4:
			bitMask = 240
		case 5:
			bitMask = 248
		case 6:
			bitMask = 252
		case 7:
			bitMask = 254
		}
		var subByte byte
		//m.matrix[0] = m.matrix[0] << dif
		for i := 1; i < len(m.matrix); i++ {
			subByte = (m.matrix[i] & bitMask) >> (8 - dif)
			m.matrix[i-1] = (m.matrix[i-1] << dif) | subByte
		}
		m.matrix[len(m.matrix)-1] = m.matrix[len(m.matrix)-1] << dif

	}
	return true
}

func (m *BoolMatrix) CountTrues() int64 {
	counter := int64(0)
	for _, i := range m.matrix {
		counter += int64(((i >> 7) & 1) + ((i >> 6) & 1) + ((i >> 5) & 1) + ((i >> 4) & 1) + ((i >> 3) & 1) + ((i >> 2) & 1) + ((i >> 1) & 1) + (i & 1))
	}
	for i := 0; i < 8-m.width*m.heigh&8; i++ {
		counter -= int64((m.matrix[len(m.matrix)-1] >> i) & 1)
	}
	return counter
}

func (m *BoolMatrix) CountTruesInLine(line int) int {
	lineStart := (line * m.width) / 8
	counter := (line * m.width) % 8
	res := 0
	for i := 0; i < m.width; i++ {
		res += int((m.matrix[lineStart+counter/8] & byte(math.Pow(2, float64(7-counter%8)))) >> (7 - (counter % 8)))
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

func (m *BoolMatrix) CheckDisbalance(disb float64) bool {
	groupsSum := make([]int, m.width)
	counter := 0
	for counter < m.width*m.heigh {
		for i := 0; i < m.width; i++ {
			switch counter % 8 {
			case 0:
				groupsSum[i] += int((m.matrix[counter/8] & 128) >> 7)
				//fmt.Println("case 0 elem value:", int((m.matrix[counter/8]&128)>>7))
			case 1:
				groupsSum[i] += int((m.matrix[counter/8] & 64) >> 6)
			case 2:
				groupsSum[i] += int((m.matrix[counter/8] & 32) >> 5)
			case 3:
				groupsSum[i] += int((m.matrix[counter/8] & 16) >> 4)
			case 4:
				groupsSum[i] += int((m.matrix[counter/8] & 8) >> 3)
			case 5:
				groupsSum[i] += int((m.matrix[counter/8] & 4) >> 2)
			case 6:
				groupsSum[i] += int((m.matrix[counter/8] & 2) >> 1)
			case 7:
				groupsSum[i] += int(m.matrix[counter/8] & 1)
			}
			//groupsSum[i] += int(m.matrix[counter/8] & byte(math.Pow(2, float64(counter%8))))
			counter++
		}
	}

	result := float64(0)
	for i := 0; i < m.width; i++ {
		result += math.Abs(float64(groupsSum[i]) - (float64(m.heigh) / float64(m.width)))

		//fmt.Println("colum sum:", groupsSum[i], " total:", result)

	}

	//fmt.Println("Disbalance:", result/float64(m.width), "sum:", result)

	if result/float64(m.width) < disb {
		return true
	} else {
		return false
	}
}

func (m *BoolMatrix) Print() {
	for i := 0; i < m.heigh; i++ {
		for j := 0; j < m.width; j++ {
			if m.GetBool(i, j) {
				fmt.Print("1 ")
			} else {
				fmt.Print("0 ")
			}
		}
		fmt.Println()
	}
}
