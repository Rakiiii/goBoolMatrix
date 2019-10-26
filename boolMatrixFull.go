package boolmatrixlib

import (
	"fmt"
	"math"
	"math/big"
)

//BoolMatrixLinear struct for holidng lineared matrix of bool
type BoolMatrixLinear struct {
	width, heigh int
	matrix       []bool
}

//Init func for initialization BoolMatrixLinear , first param is width of matrix , second param is heigh of matrix
func (b *BoolMatrixLinear) Init(w, h int) {
	b.width = w
	b.heigh = h
	b.matrix = make([]bool, w*h)
}

//SetMatrix is function for setting matrix from equls byte slice matrix needs to be initiazide before
func (b *BoolMatrixLinear) SetMatrix(newmat []byte) {
	for i, _ := range b.matrix {
		if (newmat[i/8]>>(7-(i%8)))&1 == 1 {
			b.matrix[i] = true
		} else {
			b.matrix[i] = false
		}
	}
}

//Width return matrix width
func (b *BoolMatrixLinear) Width() int {
	return b.width
}

//Heigh return matrix heigh
func (b *BoolMatrixLinear) Heigh() int {
	return b.heigh
}

//GetBool returns bool from bool matrix at positions [i,j]
func (b *BoolMatrixLinear) GetBool(i, j int) bool {
	if i > b.heigh || j > b.width {
		return false
	}

	return b.matrix[i*b.width+j]
}

//SetBool setting bool in bool matrix at position [i,j] with value
func (b *BoolMatrixLinear) SetBool(i, j int, value bool) bool {
	if i > b.heigh || j > b.width {
		return false
	}
	b.matrix[i*b.width+j] = value
	return true
}

//Print printinf bool matrix to stdout
func (b *BoolMatrixLinear) Print() {
	for i := 0; i < b.heigh; i++ {
		for j := 0; j < b.width; j++ {
			if b.GetBool(i, j) {
				fmt.Print("1 ")
			} else {
				fmt.Print("0 ")
			}
		}
		fmt.Println()
	}
}

//SetByNumber setting bool matrix equls to bit set of num
func (b *BoolMatrixLinear) SetByNumber(num *big.Int) bool {
	slice := num.Bytes()
	dif := 0

	switch {
	case len(b.matrix) < len(slice)*8:
		dif = (len(slice)*8 - len(b.matrix))
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
		for i := 1; i < len(slice); i++ {
			subByte = (slice[i] & bitMask) >> (8 - dif)
			slice[i-1] = (slice[i-1] << dif) | subByte
		}
		slice[len(slice)-1] = slice[len(slice)-1] << dif
		//return false
	case len(b.matrix) > len(slice)*8:
		slice = append(make([]byte, ((len(b.matrix)-len(slice)*8)/8)+1), slice...)
		dif = (len(slice)*8 - len(b.matrix))
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
		for i := 1; i < len(slice); i++ {
			subByte = (slice[i] & bitMask) >> (8 - dif)
			slice[i-1] = (slice[i-1] << dif) | subByte
		}
		slice[len(slice)-1] = slice[len(slice)-1] << dif
	}
	b.SetMatrix(slice)
	return true
}

//CountTrues counting amount of true in matrix
func (b *BoolMatrixLinear) CountTrues() int64 {
	counter := int64(0)
	for _, elem := range b.matrix {
		if elem {
			counter++
		}
	}
	return counter
}

//CountTruesInLine counting amount of true in line line
func (b *BoolMatrixLinear) CountTruesInLine(line int) int {
	counter := 0
	it := 0
	for it < b.width {
		if b.matrix[line*b.width+it] {
			counter++
		}
		it++
	}
	return counter
}

//return copy of matrix
func (b *BoolMatrixLinear) Copy() *BoolMatrixLinear {
	var newMatrix BoolMatrixLinear
	newMatrix.Init(b.width, b.heigh)
	for i, elem := range b.matrix {
		newMatrix.matrix[i] = elem
	}
	return &newMatrix
}

//checking the disbalance of matrix colums amount of trues
func (b *BoolMatrixLinear) CheckDisbalance(disb float64) bool {
	groupSum := make([]int, b.width)
	line := 0
	for line < b.heigh {
		pos := 0
		for pos < b.width {
			if b.matrix[line*b.width+pos] {
				groupSum[pos]++
			}
			pos++
		}
		line++
	}

	result := float64(0)
	for _, elem := range groupSum {
		result += math.Abs(float64(elem) - (float64(b.heigh) / float64(b.width)))
	}
	if result/float64(b.width) < disb {
		return true
	} else {
		return false
	}
}
