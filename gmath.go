package gmath

type Number interface {
	int | uint | uint8 | uint32 | uint64 | int8 | int16 | int32 | int64 | float32 | float64
}

// Returns the sum of args
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
func Filter[v Number](arr []v,fn func(v) bool) []v {
	res := []v{}
	for _,val := range arr {
		if fn(val) {
			res = append(res, val)
		}
	}
	return res
}