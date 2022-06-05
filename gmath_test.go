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
