package main

import (
	"fmt"
	"time"
)

func test_pointer() {
	var test_p *int
	a := 1
	test_p = &a
	fmt.Printf("value is %v, addr is %p \n", *test_p, test_p)
	// value is 1, addr is 0xc000016098
}

func base_control() {
	var a int = 2

	if b := 2; b == 1 {
		fmt.Println(1, b)
	} else if a == 2 {
		fmt.Println(2, b)
	} else {
		fmt.Println("other")
	}

	for a = 1; a < 3; a++ {
		fmt.Println(a)
	}

	var test []int = []int{1, 2, 3, 4, 5, 6}

	for index, value := range test {
		fmt.Println(index, value)
	}

	a = 2
	switch a {
	case 1:
		fmt.Println("switch", 1)
	case 2:
		fmt.Println("switch", 2)
		fmt.Println("test")
	default:
		fmt.Println("default")
	}

	for m := 1; m < 10; m++ {
		n := 1
	LOOP:
		if n <= m {
			fmt.Printf("%dx%d=%d ", n, m, m*n)
			n++
			goto LOOP
		} else {
			fmt.Println("")
		}
		n++
	}
}

func test_select() {
	chan1 := make(chan int)
	chan2 := make(chan int)

	go func() {
		time.Sleep(2 * time.Second)
		chan1 <- 1
	}()

	go func() {
		time.Sleep(1 * time.Second)
		chan2 <- 1
	}()

	// time.Sleep(1 * time.Second)
	// select {
	// case <-chan1:
	// 	fmt.Println("1")
	// case <-chan2:
	// 	fmt.Println("2")
	// default:
	// 	fmt.Println("default")
	// }
	select {
	case <-chan1:
		fmt.Println("1")
	case <-chan2:
		fmt.Println("2")
	}
}

// func

func funcname(a int, b *int) (result int) {
	defer func() {
		result++
	}()
	return 1
}

func builtIn_function() {
	chan1 := make(chan int)
	close(chan1)

	a := []int{1, 2, 3}
	a = append(a, 4)
	len_a := len(a)
	copy(a[1:], a[2:])
	println(len_a) // 底层内置的println
	println(a[1])

	b := complex(1, 2)
	print(b)

	a_memo := new(int)
	*a_memo = 1
	print(a_memo)
}

func main() {
	// test_pointer()
	// base_control()
	// test_select()
	// a := 1
	// fmt.Println(funcname(1, &a))
	builtIn_function()
}
