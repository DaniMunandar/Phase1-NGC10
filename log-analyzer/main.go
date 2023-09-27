package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {
	defer catchErr()
	result, err := analyzeLogs()
	if err != nil {
		panic(err.Error())
	}
	extractLogLevel(result)
}

func analyzeLogs() (map[string]int, error) {
	file, err := os.Open("log.txt")
	if err != nil {
		return map[string]int{}, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	result := map[string]int{}
	r, _ := regexp.Compile(`\[.+\]`) //Ini membuat ekspresi reguler yang akan digunakan untuk mencari log dalam setiap baris.
	for scanner.Scan() {
		line := scanner.Text()
		log := r.FindString(line)
		result[log]++
	}

	if err := scanner.Err(); err != nil {
		return map[string]int{}, err
	}
	return result, nil
}

func extractLogLevel(results map[string]int) {
	fmt.Println("Log Analysis Results:")
	for key, value := range results {
		fmt.Printf("%v Level\t: %v occurrences\n", key, value)
	}
}

func catchErr() {
	if r := recover(); r != nil {
		fmt.Println("Error occured:", r)
	}
}
