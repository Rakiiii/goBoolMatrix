package boolmatrixlib

import (
	"math/rand"
	"testing"
	"time"
)

func BenchmarkGetBoolTwoDimBoolArr(b *testing.B) {
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

func BenchmarkSetBoolTwoDimBoolArr(b *testing.B) {
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

func BenchmarkGetBoolBitMatrix(b *testing.B) {
	var bit BoolMatrix
	bit.Init(2, 500)
	byt := make([]byte, 125)
	rnd := rand.New(rand.NewSource(time.Now().Unix()))
	for i, _ := range byt {
		byt[i] = byte(rnd.Intn(255))
	}
	bit.SetMatrix(byt)
	for i := 0; i < b.N; i++ {
		BGetBool(&bit)
	}

}

func BenchmarkSetBoolBitMatrix(b *testing.B) {
	var bit BoolMatrix
	bit.Init(2, 500)
	byt := make([]byte, 125)
	rnd := rand.New(rand.NewSource(time.Now().Unix()))
	for i, _ := range byt {
		byt[i] = byte(rnd.Intn(255))
	}
	bit.SetMatrix(byt)
	for i := 0; i < b.N; i++ {
		BSetBool(&bit)
	}

}

func BenchmarkGetBoolFullMstrix(b *testing.B) {
	byt := make([]byte, 125)
	rnd := rand.New(rand.NewSource(time.Now().Unix()))
	for i, _ := range byt {
		byt[i] = byte(rnd.Intn(255))
	}
	var full BoolMatrixLinear
	full.Init(2, 500)
	full.SetMatrix(byt)
	for i := 0; i < b.N; i++ {
		FGetBool(&full)
	}

}

func BenchmarkSetBoolFullMstrix(b *testing.B) {
	byt := make([]byte, 125)
	rnd := rand.New(rand.NewSource(time.Now().Unix()))
	for i, _ := range byt {
		byt[i] = byte(rnd.Intn(255))
	}
	var full BoolMatrixLinear
	full.Init(2, 500)
	full.SetMatrix(byt)
	for i := 0; i < b.N; i++ {
		FSetBool(&full)
	}

}

func FGetBool(b *BoolMatrixLinear) {
	for i := 0; i < b.Heigh(); i++ {
		for j := 0; j < b.Width(); j++ {
			if b.GetBool(i, j) {
			}
		}
	}
}

func FSetBool(b *BoolMatrixLinear) {
	for i := 0; i < b.Heigh(); i++ {
		for j := 0; j < b.Width(); j++ {
			b.SetBool(i, j, true)
		}
	}
}

func BGetBool(b *BoolMatrix) {
	for i := 0; i < b.Heigh(); i++ {
		for j := 0; j < b.Width(); j++ {
			if b.GetBool(i, j) {
			}
		}
	}
}

func BSetBool(b *BoolMatrix) {
	for i := 0; i < b.Heigh(); i++ {
		for j := 0; j < b.Width(); j++ {
			b.SetBool(i, j, true)
		}
	}
}
