package gmath

import "log"

type Number interface {
	int | uint | uint8 | uint16 | uint32 | uint64 | int8 | int16 | int32 | int64 | float32 | float64
}

// SumNumber Returns the sum of args
func SumNumber[v Number](args ...v) v {
	var res v
	for _, val := range args {
		res += val
	}
	return res
}

// Min return a minimum value in args
func Min[v Number](args ...v) v {
	res := args[0]
	for _, val := range args {
		if val < res {
			res = val
		}
	}
	return res
}

// Max return a maximum value in args
func Max[v Number](args ...v) v {
	res := args[0]
	for _, val := range args {
		if val > res {
			res = val
		}
	}
	return res
}

// Sort arr according to fn(a,b),This is sort in place
// (return value of fn < 1) a will be placed before b
// (return value of fn > 1) a will be placed after b
func Sort[v Number](arr []v, fn func(v, v) v) []v {
	right := len(arr) - 1
	left := 0
	if right <= 1 {
		return arr
	}
	return quickSort(arr, left, right, fn)
}

func quickSort[v Number](arr []v, left int, right int, fn func(v, v) v) []v {
	if left < right {
		index := partition(arr, left, right, fn)
		quickSort(arr, left, index-1, fn)
		quickSort(arr, index+1, right, fn)
	}
	return arr
}

func partition[v Number](arr []v, left int, right int, compareFn func(v, v) v) int {
	tmp := arr[left]
	for left < right {
		for left < right && compareFn(arr[right], tmp) >= 0 {
			right--
		}
		arr[left] = arr[right]

		for left < right && compareFn(arr[left], tmp) <= 0 {
			left++
		}
		arr[right] = arr[left]
	}
	arr[left] = tmp
	return left
}

//Filter return a new slice filtered by fn
func Filter[v Number](arr []v, fn func(v) bool) []v {
	var res []v
	for _, val := range arr {
		if fn(val) {
			res = append(res, val)
		}
	}
	return res
}

//Reduce traverses the set of elements one by one,
//and in each step the value of the current element is added to
//the calculated result of the previous step
func Reduce[v Number](arr []v, fn func(v, v) v, initVal ...v) v {
	var res v
	if initVal != nil {
		if len(initVal) > 1 {
			log.Fatalln("\033[31m[warning]:\033[0m", "len(initVal) has to be 1")
		}
		res = fn(initVal[0], arr[0])
	} else {
		res = arr[0]
	}
	for i := 1; i < len(arr); i++ {
		res = fn(res, arr[i])
	}
	return res
}

//Reverse the position of the elements in the array and returns the array
//This method will change the original array
func Reverse[v Number](arr []v) []v {
	start := 0
	end := len(arr) - 1
	var temp v
	for start < end {
		temp = arr[start]
		arr[start] = arr[end]
		arr[end] = temp
		start++
		end--
	}
	return arr
}
