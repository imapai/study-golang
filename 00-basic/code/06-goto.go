package main

func main() {
	for num := 1; num < 10; num++ {
		println(num)
		if num > 7 {
			goto END
		}
	}
END:
	println("结束")
}
