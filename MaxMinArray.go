package main

import "fmt"

func main() {

	var n int
	fmt.Println("Введите количество элементов:")
	fmt.Scan(&n)
	if n <= 0 {
		fmt.Println("Количество элементов должно быть больше нуля.")
		return
	}

	arr := make([]float64, n)
	fmt.Println("Введите элементы массива:")

	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	min := arr[0]
	max := arr[0]
	sum := 0.0

	for _, value := range arr {
		if value < min {
			min = value
		}
		if value > max {
			max = value
		}
		sum += value
	}

	avg := sum / float64(n)
	fmt.Printf("min: %.2f\n", min)
	fmt.Printf("max: %.2f\n", max)
	fmt.Printf("avg: %.2f\n", avg)
}

//Задача 3: Поиск минимального и максимального элементов в массиве вещественных чисел
//Напишите программу, которая считывает с консоли вещественные числа и записывает их в
//массив. Затем найдите и выведите минимальный и максимальный элементы а также среднее
//значение элементов в этом массиве.
//// [0, 1, 2, 3]
//// output:
//min: 0
//max: 3
//avg: 1.5
