package main

import "fmt"

func test(s *[]string) {
	fmt.Printf("%T, %v\n", s, s)
	fmt.Printf("%T, %v\n", *s, *s)
	*s = append(*s, "test")
	*s = append(*s, "test")
	*s = append(*s, "test")
	fmt.Printf("%T, %v\n", *s, *s)
}

func test2(s []string) {
	fmt.Printf("%T, %v\n", s, s)
	s[0] = "数组是指针,所以是引用的"
	// slice的数据是放在数组指针的，所以刚传进来的数组地址是按值传的 和外面值一样，
	// 所以改变数组元素也把外面的值变了
	s = append(s, "test")
	s = append(s, "test")
	s[1] = "扩容改变了地址，所以这里没改变外面"
	fmt.Printf("%T, %v\n", s, s)
}

func main() {
	tt := []string{"q", "qq"}
	// test(&tt)
	test2(tt)
	fmt.Print(tt, &tt)
}

/*
[]string, [q qq]
[]string, [数组是指针,所以是引用的 扩容改变了地址，所以这里没改变外面 test test]                                                                              ]
[数组是指针,所以是引用的 qq] &[数组是指针,所以是引用的 qq]
*/
