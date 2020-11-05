package main

func main() {
	var i int
	for {
		if i > 5 {
			break
		}
		println(i)
		i++
	}
	println(i) // print 6
}
