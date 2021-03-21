package main

type rect struct {
	width, height int
}

func (r rect) area() int {
	return r.width * r.height
}
func (r *rect) perimeter() int {
	return (r.width + r.height) * 2
}

func main() {
	r := rect{height: 2, width: 3}
	println(r.area())
	println(r.perimeter())

	rp := &r
	println(rp.area())
	println(rp.perimeter())
}
