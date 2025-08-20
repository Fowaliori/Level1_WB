package main

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
