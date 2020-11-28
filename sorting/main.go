package main

import "fmt"

func main() {
	a := [7]int{4, 9, 7, 5, 8, 9, 3}
	swap := false
	count := 0
	if swap == true {
		for i := 0; i < 10; i++ {
			if a[i] > a[i+1] {
				temp := a[1]
				a[i] = a[i+1]
				a[i+1] = temp
				count += 1
				fmt.Println("[", a[i], ",", a[i+1], "]", "->", a)
			}
		}
	}
}

//!done
