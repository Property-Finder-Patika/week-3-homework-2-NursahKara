package main

import "fmt"

type rectangle struct {
	length  float64
	breadth float64

	geometry struct {
		area      float64
		perimeter float64
	}
}

func area(length float64, breadth float64) float64 {

	return length * breadth
}
func primeter(length float64, breadth float64) float64 {
	return 2 * (length + breadth)
}
func main() {
	var rect rectangle
	fmt.Scanf("%f %f", &rect.length, &rect.breadth)

	rect.geometry.area = area(rect.length, rect.breadth)
	rect.geometry.perimeter = primeter(rect.length, rect.breadth)

	fmt.Println(rect)
	fmt.Println("Area:\t", rect.geometry.area)
	fmt.Println("Perimeter:", rect.geometry.perimeter)
}
