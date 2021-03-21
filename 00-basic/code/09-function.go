package main

func main() {
	println(add(10, 12))
	println(multipleAdd(1, 2))
	println(multipleAdd(1, 2, 3))
}

func add(x, y int) int {
	return x + y
}

func multipleAdd(nums ...int) int {
	result := 0
	for _, n := range nums {
		result += n
	}
	return result
}
