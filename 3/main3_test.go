package main

import "testing"

func TestMap(t *testing.T) {
	res := NewStringIntMap()

	res.Add("Gopher", 10)
	val, ok := res.Get("Gopher")
	if !ok || val != 10 {
		t.Error("Ошибка: не удалось добавить или получить Gopher")
	}

	if !res.Exists("Gopher") {
		t.Error("Ошибка: должен был вернуть true")
	}

	copyData := res.Copy()
	if copyData["Gopher"] != 10 {
		t.Error("Ошибка: в копии нет данных из оригинала")
	}
	res.Add("Gopher", 999)
	if copyData["Gopher"] == 999 {
		t.Error("Ошибка: копия изменилась вместе в оригиналом (это не копия!)")
	}

	res.Remove("Gopher")
	if res.Exists("Gopher") {
		t.Error("Ошибка: Gopher все еще существует после удаления")
	}

	val, ok = res.Get("Gopher")
	if ok || val != 0 {
		t.Error("Ошибка: Get должен возвращать (0, false) для удаленного ключа")
	}
}
