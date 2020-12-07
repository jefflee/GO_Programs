package main
import (
	"fmt"
	"reflect"
	"runtime"
	"strings"
	"time"
)


func main() {

	ourTitle := "the go standard library"
	timed := MakeTimedFunction(properTitle).(func(string) string)
	newTitle := timed(ourTitle)
	fmt.Println(newTitle)
	timedToo := MakeTimedFunction(doubleOurNumber).(func(int) int)
	fmt.Println(timedToo(2))
}

func MakeTimedFunction(f interface{}) interface{} {
	rf := reflect.TypeOf(f)
	if rf.Kind() != reflect.Func {
		panic("expects a function")
	}
	vf := reflect.ValueOf(f)
	wrapperF := reflect.MakeFunc(rf, func(in []reflect.Value) []reflect.Value {
		start := time.Now()
		out := vf.Call(in)
		end := time.Now()
		fmt.Printf("calling %s took %v\n", runtime.FuncForPC(vf.Pointer()).Name(), end.Sub(start))
		return out
	})
	return wrapperF.Interface()
}

func properTitle(input string) string {
	// from: https://golangcookbook.com/chapters/strings/title/
	words := strings.Fields(input)
	smallwords := " a an on the to "

	for index, word := range words {
		if strings.Contains(smallwords, " "+word+" ") {
			words[index] = word
		} else {
			words[index] = strings.Title(word)
		}
	}
	return strings.Join(words, " ")
}

func doubleOurNumber(a int) int {f
	time.Sleep(1 * time.Second)
	return a * 2
}