package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

func isAnagram(word1, word2 string) bool {
	// Menghilangkan spasi dan mengubah huruf menjadi huruf kecil agar tidak sensitif terhadap kapitalisasi
	word1 = strings.ToLower(strings.ReplaceAll(word1, " ", ""))
	word2 = strings.ToLower(strings.ReplaceAll(word2, " ", ""))

	// Mengonversi string menjadi slice rune untuk pengurutan
	runeWord1 := []rune(word1)
	runeWord2 := []rune(word2)

	// Mengurutkan kedua kata
	sort.Slice(runeWord1, func(i, j int) bool {
		return runeWord1[i] < runeWord1[j]
	})
	sort.Slice(runeWord2, func(i, j int) bool {
		return runeWord2[i] < runeWord2[j]
	})

	// Membandingkan kata-kata yang telah diurutkan
	return string(runeWord1) == string(runeWord2)
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Error:", r)
		}
	}()

	if len(os.Args) != 3 {
		panic("Jumlah argumen harus 2")
	}

	word1 := os.Args[1]
	word2 := os.Args[2]

	// Validasi panjang kata
	if len(word1) > 10 || len(word2) > 10 {
		panic("Input tidak valid: Kata melebihi 10 karakter")
	}

	// Validasi karakter yang diperbolehkan
	invalidChars := "!@#$%^&*()-+[]{};:'\"<>?/.,"
	for _, char := range invalidChars {
		if strings.Contains(word1, string(char)) || strings.Contains(word2, string(char)) {
			panic("Input tidak valid: Kata mengandung karakter yang tidak diperbolehkan")
		}
	}

	if isAnagram(word1, word2) {
		fmt.Printf("[%s] & [%s] merupakan anagram\n", word1, word2)
	} else {
		fmt.Printf("[%s] & [%s] bukan merupakan anagram\n", word1, word2)
	}
}
