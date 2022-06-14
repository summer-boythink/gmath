package gmath

import (
	"reflect"
	"testing"
)

func TestSumNumber(t *testing.T) {
	a, b, c := 11, 21, -2
	res := SumNumber(a, b, c)
	if res != 30 {
		t.Errorf("want 30,get %d", res)
	}
}

func TestMin(t *testing.T) {
	a := []int32{11, 22, 9, 2}
	if res := Min(a...); res != 2 {
		t.Errorf("want 2,get %d", res)
	}
}

func TestMax(t *testing.T) {
	a := []int32{11, 22, 9, 2}
	if res := Max(a...); res != 22 {
		t.Errorf("want 22,get %d", res)
	}
}

func TestSort(t *testing.T) {
	arr := []int32{11, 22, 9, 2}
	Sort(arr, func(a, b int32) int32 {
		return a - b
	})
	if !reflect.DeepEqual(arr, []int32{2, 9, 11, 22}) {
		t.Error("want [2,9,11,22],but get", arr)
	}
}

func TestFilter(t *testing.T) {
	arr := []float64{11.00, 12.02, 13}
	fn := func(val float64) bool {
		return val > 12
	}
	res := Filter(arr, fn)
	if !reflect.DeepEqual(res, []float64{12.02, 13}) {
		t.Error("want [12.02 13],but get", res)
	}
}

func TestReduce(t *testing.T) {
	arr := []uint16{11, 22, 134}
	fn := func(a, b uint16) uint16 {
		return a + b
	}
	res1 := Reduce(arr, fn)
	want1 := SumNumber(arr...)
	if res1 != want1 {
		t.Errorf(" want %d,but get %d", want1, res1)
	}

	res2 := Reduce(arr, fn, 122)
	want2 := SumNumber(arr...) + 122
	if res2 != want2 {
		t.Errorf(" want %d,but get %d", want2, res2)
	}
}

func TestReverse(t *testing.T) {
	arr := []int{1, 23, 5, 2, 12}
	Reverse(arr)
	if !reflect.DeepEqual(arr, []int{12, 2, 5, 23, 1}) {
		t.Error("want []int{12,2,5,23,1},but get", arr)
	}
}
