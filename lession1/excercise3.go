package lession1

import "fmt"

func SliceHandler() {
	sl := []float64{2, 3, 4, 15, 7}
	fmt.Println("Tong cac phan tu thuoc slice: ", sumOfSlice(sl))
	fmt.Println("Trung binh cong cac phan tu thuoc slice: ", averageSlice(sl))
	min, max := findMinMax(sl)
	fmt.Println("Gia tri lon nhat cua slice la ", max, "va nho nhat: ", min)
}

func sumOfSlice(sl []float64) float64 {
	var sum float64 = 0
	for _, e := range sl {
		sum += e
	}
	return sum
}

func averageSlice(sl []float64) float64 {
	sum := sumOfSlice(sl)
	return sum / float64(len(sl))
}

func findMinMax(sl []float64) (min, max float64) {
	min = sl[0]
	max = sl[0]
	for _, e := range sl {
		if e <= min {
			min = e
		}
		if e >= max {
			max = e
		}
	}
	return
}
