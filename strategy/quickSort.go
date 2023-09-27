package strategy

type QuickSort struct{}

func (qs *QuickSort) Sort(arr []int) []int {
	return quickSort(arr)
}

func quickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	pivot := arr[0]
	var left, right []int
	for _, nb := range arr[1:] {
		if nb <= pivot {
			left = append(left, nb)
		} else {
			right = append(right, nb)
		}
	}

	left = quickSort(left)
	right = quickSort(right)

	return append(append(left, pivot), right...)
}
