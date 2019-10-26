package boolmatrixlib

import (
	"fmt"
	"math/big"
	"testing"
)

func TestGetBoolF(t *testing.T) {
	c := make([]byte, 5)
	c[0] = 128
	c[1] = 64
	c[2] = 32
	c[3] = 16
	c[4] = 8
	var b BoolMatrixLinear
	b.Init(8, 5)
	b.SetMatrix(c)
	if !b.GetBool(0, 0) && !b.GetBool(1, 1) && b.GetBool(2, 3) {
		t.Error("Wrong get bool")
	}
}

func TestSetBoolF(t *testing.T) {
	c := make([]byte, 5)
	c[0] = 128
	c[1] = 64
	c[2] = 32
	c[3] = 16
	c[4] = 8
	var b BoolMatrixLinear
	b.Init(8, 5)
	b.SetMatrix(c)
	b.SetBool(1, 1, false)
	b.SetBool(4, 6, true)
	if b.GetBool(1, 1) {
		t.Error("Wrong sey at 1,1")
		fmt.Println("Matrix:")
		fmt.Println()
		b.Print()
	}
	if !b.GetBool(4, 6) {
		t.Error("Wrong sey at 4,6")
		b.Print()
	}
}

func TestSetByNumberF(t *testing.T) {
	var b BoolMatrixLinear
	b.Init(8, 2)
	b.SetByNumber(big.NewInt(int64(65280)))
	var d BoolMatrixLinear
	d.Init(8, 2)
	c := make([]byte, 2)
	c[0] = 255
	c[1] = 0
	d.SetMatrix(c)
	flag := false

	for i := 0; i < 2; i++ {
		for j := 0; j < 8; j++ {
			if b.GetBool(i, j) != d.GetBool(i, j) {
				t.Error("Differenc in at pos ", i, " ", j)
				flag = true
			}
		}
	}

	if flag {
		fmt.Println("Matrix by number 1:")
		fmt.Println()

		b.Print()
		fmt.Println("Matrix by set matrix 1:")
		fmt.Println()

		d.Print()
	}

	var d2 BoolMatrixLinear
	d2.Init(5, 3)
	d2.SetByNumber(big.NewInt(int64(255)))

	c2 := make([]byte, 2)
	c2[0] = 1
	c2[1] = 255
	var b2 BoolMatrixLinear
	b2.Init(5, 3)
	b2.SetMatrix(c2)
	flag = false

	for i := 0; i < 2; i++ {
		for j := 0; j < 8; j++ {
			if b2.GetBool(i, j) != d2.GetBool(i, j) {
				t.Error("Differenc in at pos ", i, " ", j)
				flag = true
			}
		}
	}

	if flag {
		fmt.Println("Matrix by number 2:")
		fmt.Println()

		d2.Print()
		fmt.Println("Matrix by set matrix 2:")
		fmt.Println()

		b2.Print()

		fmt.Println("expected:")
		fmt.Println()
		var d3 BoolMatrix
		d3.Init(5, 3)
		d3.SetByNumber(big.NewInt(int64(255)))
		d3.Print()
	}

	//var m BoolMatrixLinear
	//m.Init(2, 5)
	//m.SetByNumber(big.NewInt(int64(1016)))
	//m.Print()

}

func TestCountTrueF(t *testing.T) {
	var d BoolMatrixLinear
	d.Init(5, 3)
	d.SetByNumber(big.NewInt(int64(255)))
	if d.CountTrues() != 8 {
		t.Error("Wrong Amount of trues:", d.CountTrues())
		fmt.Println("Matrix:")
		d.Print()
	}
}

func TestCountTrueInLineF(t *testing.T) {
	var b BoolMatrixLinear
	b.Init(8, 2)
	b.SetByNumber(big.NewInt(int64(65280)))
	if b.CountTruesInLine(0) != 8 {
		t.Error("Wrong amount of trues in line 0:", b.CountTruesInLine(0))
		fmt.Println("Matrix:")
		b.Print()
	}

	if b.CountTruesInLine(1) != 0 {
		t.Error("Wrong amount of trues in line 1:", b.CountTruesInLine(1))
		fmt.Println("Matrix:")
		b.Print()
	}

	b.SetByNumber(big.NewInt(int64(49040)))
	if b.CountTruesInLine(0) != 7 {
		t.Error("Wrong amount of trues in line 0,matrix 2:", b.CountTruesInLine(0))
		fmt.Println("Matrix:")
		b.Print()

	}

	var d BoolMatrixLinear
	d.Init(5, 3)
	d.SetByNumber(big.NewInt(int64(255)))
	if d.CountTruesInLine(2) != 5 {
		t.Error("Wrong amount of true in cust matrix,line 0:", d.CountTruesInLine(0))
		//fmt.Println("Bigr int:")
		fmt.Println("Matrix:")
		d.Print()

	}
	if d.CountTruesInLine(1) != 3 {
		t.Error("Wrong amount of true in cust matrix,line 1:", d.CountTruesInLine(1))
		fmt.Println("Matrix:")
		d.Print()

	}

}

func TestCheckDisbalanceF(t *testing.T) {
	var b BoolMatrixLinear
	b.Init(3, 4)
	b.SetByNumber(big.NewInt(int64(2217)))

	if !b.CheckDisbalance(0.7) {
		t.Error("Wrong disbalance")
		fmt.Println("matrix 0,7 with wrong disb")
		b.Print()
	}

	fmt.Println()
	var d BoolMatrixLinear
	d.Init(2, 5)
	d.SetByNumber(big.NewInt(int64(682)))
	//d.Print()
	if !d.CheckDisbalance(2.6) || d.CheckDisbalance(1) {
		t.Error("Wrong CheckDisbalance")
		fmt.Println("matrix with wrong disb:")
		d.Print()
	}
}
