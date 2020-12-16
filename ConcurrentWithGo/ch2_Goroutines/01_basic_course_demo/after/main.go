package main

import (
	"fmt"
	"math/rand"
	"time"

	"concurrent/books"
)

var cache = map[int]books.Book{}
var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

func main() {
	for i := 0; i < 10; i++ {
		id := rnd.Intn(10) + 1
		if b, ok := queryCache(id); ok {
			fmt.Println("from cache")
			fmt.Println(b)
			continue
		}
		if b, ok := queryDatabase(id); ok {
			fmt.Println("from database")
			cache[id] = b
			fmt.Println(b)
			continue
		}
		fmt.Printf("Book not found id: '%v'", id)
	}
}

func queryCache(id int) (books.Book, bool) {
	b, ok := cache[id]
	return b, ok
}

func queryDatabase(id int) (books.Book, bool) {
	time.Sleep(300 * time.Millisecond)
	for _, b := range books.Books {
		if b.ID == id {
			return b, true
		}
	}

	return books.Book{}, false
}
