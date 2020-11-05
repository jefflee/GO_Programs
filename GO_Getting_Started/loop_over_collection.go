package main

func main() {
	slice := []int{1, 2, 3}
	for i, v := range slice {
		println(i, v)
	}

	portMap := map[string]int{"http": 80, "https": 443}
	for k, v := range portMap {
		println(k, v)
	}

	// if we don't need value
	for k := range portMap {
		println(k)
	}

	// if we don't need key
	for _, v := range portMap {
		println(v)
	}
}
