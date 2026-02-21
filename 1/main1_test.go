package main

import "testing"

func TestGetType(t *testing.T) {
	if GetType(42) != "int" {
		t.Error("Ошибка: тип должен быть int")
	}
	if GetType(3.14) != "float64" {
		t.Error("Ошибка: тип должен быть float64")
	}
	if GetType("Golang") != "string" {
		t.Error("Ошибка: тип должен быть string")
	}
	if GetType(true) != "bool" {
		t.Error("Ошибка: тип должен быть bool")
	}
	if GetType(1+2i) != "complex64" {
		t.Error("Ошибка: тип должен быть complex64")
	}
}

func TestConcat(t *testing.T) {
	result := Concat("A", 1, true)
	if result != "A1true" {
		t.Errorf("Ожидали A1true, но получили %s", result)
	}
}

func TestHashSHA256(t *testing.T) {
	result := HashSHA256([]rune("test"), "salt")
	if len(result) != 64 {
		t.Error("Хеш должен быть длиной 64 символа")
	}
}
