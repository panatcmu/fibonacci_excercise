package main

import "fmt"

func fibonacci() func() int {
	var j, k int = 0 , 1
	return func() int {
			j , k = k , j + k
			return k - j
	}
}


func main() {
	f := fibonacci()
	for i := 0; i<10; i++ {
		fmt.Println(f())
	}
}
