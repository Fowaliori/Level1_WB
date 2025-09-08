package main

/*
Дана структура Human (с произвольным набором полей и методов).

Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).

Подсказка: используйте композицию (embedded struct), чтобы Action имел все методы Human.
*/

type human struct {
	age int
}

type action struct {
	human
}

func (h *human) grow() int {
	h.age = h.age + 1
	return h.age
}

func main() {
	bair := action{human{25}}
	bair.grow()
	println(bair.age)
}
