package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Terjadi kesalahan", r)
		}
	}()

	fmt.Println("Pilih operasi aritmatika")
	fmt.Println("> Penjumlahan (+) a")
	fmt.Println("> Pengurangan (-) b")
	fmt.Println("> Perkalian (*) c")
	fmt.Println("> Pembagian (/) d")

	reader := bufio.NewReader(os.Stdin)
	operation, _ := reader.ReadString('\n')
	operation = strings.TrimSpace(operation)

	switch operation {
	case "a":
		fmt.Println("Masukan angka pertama:")
		num1 := readNumber()
		fmt.Println("Masukan angka kedua:")
		num2 := readNumber()
		result := num1 + num2
		fmt.Printf("Hasil Penjumlahan %d dan %d adalah %d\n", num1, num2, result)
	case "b":
		fmt.Println("Masukan angka pertama:")
		num1 := readNumber()
		fmt.Println("Masukan angka kedua:")
		num2 := readNumber()
		result := num1 - num2
		fmt.Printf("Hasil Pengurangan %d dan %d adalah %d\n", num1, num2, result)
	case "c":
		fmt.Println("Masukan angka pertama:")
		num1 := readNumber()
		fmt.Println("Masukan angka kedua:")
		num2 := readNumber()
		result := num1 * num2
		fmt.Printf("Hasil Perkalian %d dan %d adalah %d\n", num1, num2, result)
	case "d":
		fmt.Println("Masukan angka pertama:")
		num1 := readNumber()
		fmt.Println("Masukan angka kedua:")
		num2 := readNumber()
		if num2 == 0 {
			panic("Pembagian oleh nol tidak diizinkan")
		}
		result := float64(num1) / float64(num2)
		fmt.Printf("Hasil Pembagian %d dan %d adalah %.2f\n", num1, num2, result)
	default:
		panic("Operasi yang di pilih tidak valid.")
	}
}

func readNumber() int {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	num, err := strconv.Atoi(input)
	if err != nil {
		panic("Masukan harus berupa angka.")
	}
	return num
}
