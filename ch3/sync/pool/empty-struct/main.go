package main

func main() {
	a := struct{}{}
	b := struct{}{}
	println(&a, &b, &a == &b)

	c := new(struct{})
	d := new(struct{})
	println(c, d, c == d)
}