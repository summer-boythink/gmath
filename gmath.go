package gmath

type Number interface {
	int | uint | uint8 | uint32 | uint64 | int8 | int16 | int32 | int64 | float32 | float64
}

func SumNumber[v Number](args ...v) v {
	var res v
	for _, val := range args {
		res += val
	}
	return res
}

func Min[v Number](args ...v) v {
	res := args[0]
	for _, val := range args {
		if val < res {
			res = val
		}
	}
	return res
}

func Max[v Number](args ...v) v {
	res := args[0]
	for _, val := range args {
		if val > res {
			res = val
		}
	}
	return res
}

func Sort[v Number](arr []v, fn func(v, v) v) []v {
	right := len(arr) - 1
	left := 0
	if right <= 1 {
		return arr
	}
	// quickSort
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
