package boolmatrixlib

import "math/big"

//IBoolMatrix interface implimenting all functions to work with different bool matrix
type IBoolMatrix interface {
	Init(w, h int)
	SetMatrix(newmat []byte)
	Width() int
	Heigh() int
	GetBool(i, j int) bool
	SetBool(i, j int, value bool) bool
	Print()
	SetByNumber(num *big.Int) bool
	CountTrues() int64
	CountTruesInLine(line int) int
	CopyIBoolMatrix() IBoolMatrix
	CheckDisbalance(disb float64) bool
}
