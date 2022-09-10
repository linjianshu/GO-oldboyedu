package main

import (
	"fmt"
)

type a struct {
	val  int
	next *a
}

var b bool

func main() {
	var a1, a2, a3, a4, a5 a
	a1 = a{
		val:  1,
		next: &a2,
	}

	a2 = a{
		val:  2,
		next: &a3,
	}

	a3 = a{
		val:  3,
		next: &a4,
	}

	a4 = a{
		val:  4,
		next: &a5,
	}

	a5 = a{
		val:  5,
		next: nil,
	}
	judge(a1, a1)
	fmt.Println(b)
}

func judge(a11, a12 a) {
	next := *a11.next
	nextnext := *a12.next.next
	if nextnext.next == nil {
		b = false
		return
	}
	if next == nextnext {
		b = true
		return
	}
	judge(next, nextnext)
}
