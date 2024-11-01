package lession1

import "fmt"

func AreaAndPerimeterOfRectangle() {
	fmt.Println("----------Excercise 1----------")
	var x, y float64
	fmt.Print("Nhap vao chieu dai va rong hcn: ")
	fmt.Scanf("%f %f", &x, &y)
	fmt.Println("Dien tich hinh chu nhat: ", areaOfRectangle(x, y))
	fmt.Println("Chu vi hinh chu nhat: ", perimeterOfRectangle(x, y))
}
func areaOfRectangle(x, y float64) float64 {
	return x * y
}

func perimeterOfRectangle(x, y float64) float64 {
	return 2 * (x + y)
}
