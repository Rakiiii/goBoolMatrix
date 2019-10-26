package boolmatrixlib

import (
	"math/rand"
	"testing"
	"time"
)

var byt []byte = *initByt()
var xrnd []int = *initX()
var yrnd []int = *initY()

func BenchmarkGetBoolTwoDimBoolArrAllElems(b *testing.B) {
	m := make([][]bool, 500)
	for i, _ := range m {
		m[i] = make([]bool, 2)
	}
	for k := 0; k < b.N; k++ {
		for i := 0; i < len(m); i++ {
			for j := 0; j < len(m[i]); j++ {
				if m[i][j] {

				}
			}
		}
	}
}

func BenchmarkSetBoolTwoDimBoolArrAllElems(b *testing.B) {
	m := make([][]bool, 500)
	for i, _ := range m {
		m[i] = make([]bool, 2)
	}
	for k := 0; k < b.N; k++ {
		for i := 0; i < len(m); i++ {
			for j := 0; j < len(m[i]); j++ {
				m[i][j] = true
			}
		}
	}
}

func BenchmarkGetBoolBitMatrixAllElems(b *testing.B) {
	var bit BoolMatrix
	bit.Init(2, 500)
	bit.SetMatrix(byt)
	for i := 0; i < b.N; i++ {
		BGetBoolAllElems(&bit)
	}

}

func BenchmarkSetBoolBitMatrixAllElems(b *testing.B) {
	var bit BoolMatrix
	bit.Init(2, 500)
	bit.SetMatrix(byt)
	for i := 0; i < b.N; i++ {
		BSetBoolAllElems(&bit)
	}

}

func BenchmarkGetBoolFullMstrixAllElems(b *testing.B) {
	var full BoolMatrixLinear
	full.Init(2, 500)
	full.SetMatrix(byt)
	for i := 0; i < b.N; i++ {
		FGetBoolAllElems(&full)
	}

}

func BenchmarkSetBoolFullMstrixAllElems(b *testing.B) {
	var full BoolMatrixLinear
	full.Init(2, 500)
	full.SetMatrix(byt)
	for i := 0; i < b.N; i++ {
		FSetBoolAllElems(&full)
	}

}

func BenchmarkGetBoolTwoDimBoolArrRandomAdress(b *testing.B) {
	m := make([][]bool, 500)
	for i, _ := range m {
		m[i] = make([]bool, 2)
	}
	for k := 0; k < b.N; k++ {
		if m[yrnd[k%1000]][xrnd[k%1000]] {

		}
	}
}

func BenchmarkSetBoolTwoDimBoolArrRandomAdress(b *testing.B) {
	m := make([][]bool, 500)
	for i, _ := range m {
		m[i] = make([]bool, 2)
	}
	for k := 0; k < b.N; k++ {
		m[yrnd[k%1000]][xrnd[k%1000]] = true

	}
}

func BenchmarkGetBoolBitMatrixRandomAdress(b *testing.B) {
	var bit BoolMatrix
	bit.Init(2, 500)
	bit.SetMatrix(byt)
	for i := 0; i < b.N; i++ {
		bit.GetBool(yrnd[i%1000], xrnd[i%1000])
	}
}

func BenchmarkSetBoolBitMatrixRandomAdress(b *testing.B) {
	var bit BoolMatrix
	bit.Init(2, 500)
	bit.SetMatrix(byt)
	for i := 0; i < b.N; i++ {
		bit.SetBool(yrnd[i%1000], xrnd[i%1000], true)
	}
}

func BenchmarkGetBoolFullMatrixRandomAdress(b *testing.B) {
	var full BoolMatrixLinear
	full.Init(2, 500)
	full.SetMatrix(byt)
	for i := 0; i < b.N; i++ {
		full.GetBool(yrnd[i%1000], xrnd[i%1000])
	}
}

func BenchmarkSetBoolFullMatrixRandomAdress(b *testing.B) {
	var full BoolMatrixLinear
	full.Init(2, 500)
	full.SetMatrix(byt)
	for i := 0; i < b.N; i++ {
		full.SetBool(yrnd[i%1000], xrnd[i%1000], true)
	}
}

func BenchmarkCountTruesBitMatrix(b *testing.B) {
	var bit BoolMatrix
	bit.Init(2, 500)
	bit.SetMatrix(byt)
	for i := 0; i < b.N; i++ {
		bit.CountTrues()
	}
}

func BenchmarkCountTruesFullMatrix(b *testing.B) {
	var full BoolMatrixLinear
	full.Init(2, 500)
	full.SetMatrix(byt)
	for i := 0; i < b.N; i++ {
		full.CountTrues()
	}
}

func BenchmarkCountTruesInLineBitMatrix(b *testing.B) {
	var bit BoolMatrix
	bit.Init(2, 500)
	bit.SetMatrix(byt)
	for i := 0; i < b.N; i++ {
		bit.CountTruesInLine(yrnd[i%1000])
	}
}

func BenchmarkCountTruesInLineFullMatrix(b *testing.B) {
	var full BoolMatrixLinear
	full.Init(2, 500)
	full.SetMatrix(byt)
	for i := 0; i < b.N; i++ {
		full.CountTruesInLine(yrnd[i%1000])
	}
}

func BenchmarkCheckDisbalanceBitMatrix(b *testing.B) {
	var bit BoolMatrix
	bit.Init(2, 500)
	bit.SetMatrix(byt)
	for i := 0; i < b.N; i++ {
		bit.CheckDisbalance(float64(10))
	}
}

func BenchmarkCheckDisbalanceFullMatrix(b *testing.B) {
	var bit BoolMatrixLinear
	bit.Init(2, 500)
	bit.SetMatrix(byt)
	for i := 0; i < b.N; i++ {
		bit.CheckDisbalance(float64(10))
	}
}

func FGetBoolAllElems(b *BoolMatrixLinear) {
	for i := 0; i < b.Heigh(); i++ {
		for j := 0; j < b.Width(); j++ {
			if b.GetBool(i, j) {
			}
		}
	}
}

func FSetBoolAllElems(b *BoolMatrixLinear) {
	for i := 0; i < b.Heigh(); i++ {
		for j := 0; j < b.Width(); j++ {
			b.SetBool(i, j, true)
		}
	}
}

func BGetBoolAllElems(b *BoolMatrix) {
	for i := 0; i < b.Heigh(); i++ {
		for j := 0; j < b.Width(); j++ {
			if b.GetBool(i, j) {
			}
		}
	}
}

func BSetBoolAllElems(b *BoolMatrix) {
	for i := 0; i < b.Heigh(); i++ {
		for j := 0; j < b.Width(); j++ {
			b.SetBool(i, j, true)
		}
	}
}

func initByt() *[]byte {
	byt := make([]byte, 125)
	rnd := rand.New(rand.NewSource(time.Now().Unix()))
	for i, _ := range byt {
		byt[i] = byte(rnd.Intn(255))
	}
	return &byt
}

func initX() *[]int {
	x := make([]int, 1000)
	rnd := rand.New(rand.NewSource(time.Now().Unix()))
	for i, _ := range x {
		x[i] = rnd.Intn(2)
	}
	return &x
}

func initY() *[]int {
	y := make([]int, 1000)
	rnd := rand.New(rand.NewSource(time.Now().Unix()))
	for i, _ := range y {
		y[i] = rnd.Intn(500)
	}
	return &y
}
