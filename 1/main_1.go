package main

import (
	"crypto/sha256"
	"fmt"
)

func GetType(v interface{}) string {
	return fmt.Sprintf("%T", v)
}

func Concat(args ...interface{}) string {
	var res string
	for _, arg := range args {
		res += fmt.Sprintf("%v", arg)
	}
	return res
}

func HashSHA256(runes []rune, salt string) string {
	middle := len(runes) / 2
	newString := string(runes[:middle]) + salt + string(runes[middle:])

	hash := sha256.Sum256([]byte(newString))
	return fmt.Sprintf("%x", hash)
}

func main() {
	var numDecimal int = 42
	var numOctal int = 052
	var numHexadecimal int = 0x2A
	var pi float64 = 3.14
	var name string = "Golang"
	var isActive bool = true
	var complexNum complex64 = 1 + 2i

	data := []interface{}{numDecimal, numOctal, numHexadecimal, pi, name, isActive, complexNum}

	for _, v := range data {
		fmt.Printf("Значение: %v Тип: %s\n", v, GetType(v))
	}

	fullString := Concat(data...)
	runeSlice := []rune(fullString)

	salt := "go-2024"
	readyHash := HashSHA256(runeSlice, salt)

	fmt.Printf("Объединенная строка: %s\n", fullString)
	fmt.Printf("Срез рун: %s\n", string(runeSlice))
	fmt.Printf("SHA256 (с солью): %s\n", readyHash)
}
