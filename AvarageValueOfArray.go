package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Print("Введите количество элементов массива: ")
	fmt.Scan(&n)
	var arr [10]float64
	fmt.Println("Введите элементы массива:")
	for i := 0; i < n; i++ {
		fmt.Printf("Элемент %d: ", i+1)
		fmt.Scan(&arr[i])
	}
	sum := 0.0
	for i := 0; i < n; i++ {
		sum += arr[i]
	}
	average := sum / float64(n)
	fmt.Println("Ваш массив: ", arr)
	fmt.Printf("Сумма элементов массива: %.2f\n", sum)
	fmt.Printf("Среднее значение элементов массива: %.2f\n", average)
}

//Задача 1: Найти среднее значение элементов массива дробных чисел
//Напишите программу, которая запрашивает у пользователя ввод дробных чисел (float64)
//для заполнения массива и затем вычисляет и выводит среднее значение элементов массива.
//summa / count
//
