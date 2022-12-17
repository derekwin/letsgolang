package main

import (
	"fmt"
	"time"
)

// 变量初始化
var a bool = true
var b int = 64
var c string = "test"

func test_data() {
	var d int
	e := "hello"
	var f bool
	const g = 12

	fmt.Printf("a type : %T, a value : %v\n", a, a)
	fmt.Printf("b type : %T, b value : %v\n", b, b)
	fmt.Printf("c type : %T, c value : %v\n", c, c)
	fmt.Printf("d type : %T, d value : %v\n", d, d)
	fmt.Printf("e type : %T, e value : %v\n", e, e)
	fmt.Printf("f type : %T, f value : %v\n", f, f)
	fmt.Printf("g type : %T, g value : %v\n", g, g)

	// 位运算
	h := 1
	i := 0
	j := h | i
	fmt.Println(j)

	/*
		result:
		a type : bool, a value : true
		b type : int, b value : 64
		c type : string, c value : test
		d type : int, d value : 0
		e type : string, e value : hello
		f type : bool, f value : false
		g type : int, g value : 12
	*/
}

func test_array() {
	// 数组
	var tt [3]int = [3]int{1, 1}
	tt_1 := [2]string{"string1", "string2"}
	tt_2 := [...]float32{13: 5.0} //数字索引
	fmt.Println(tt[2], tt_1[1], tt_2[11], len(tt_2))
	// reuslt: 0 string2 0 14
}

func test_slice() {
	// 切片初始化与扩容变化
	var ss [10]int = [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s := ss[6:9]
	fmt.Println(s, &s[0], &ss[6])
	// [6 7 8] 0xc0000123a0 0xc0000123a0
	s[0] = 3
	fmt.Println(s, s[0], ss[6])
	// [3 7 8] 3 3
	fmt.Println(cap(s))
	// 4
	s = append(s, 11)
	s = append(s, 12)
	fmt.Println(cap(s))
	// 8
	fmt.Println(s, &s[0], &ss[6])
	// [3 7 8 11 12] 0xc000014440 0xc0000123a0
	s[0] = 4
	fmt.Println(s, s[0], ss[6])
	//[4 7 8 11 12] 4 3

	s_1 := []int{1, 2, 3, 4}
	fmt.Println(s_1)
	//result: [1 2 3 4]

	s_2 := s_1[1:2]
	fmt.Println(s_2)
	//result: [2]

	//扩容机制引起的底层数组值覆盖现象
	s_3 := []int{1}
	s_3 = append(s_3, 2)
	s_3 = append(s_3, 3)
	fmt.Println(s_3, &s_3[0], cap(s_3), len(s_3))
	s_4 := append(s_3, 6)
	s_5 := append(s_3, 7)
	fmt.Println(s_4, &s_4[0], cap(s_4), len(s_4))
	fmt.Println(s_5, &s_5[0], cap(s_5), len(s_5))
	/*
		[1 2 3] 0xc000016260 4 3
		[1 2 3 7] 0xc000016260 4 4
		[1 2 3 7] 0xc000016260 4 4
	*/

	// s_3 := []int{1, 2, 3}
	// fmt.Println(s_3, &s_3[0], cap(s_3), len(s_3))
	// s_4 := append(s_3, 6)
	// s_5 := append(s_3, 7)
	// fmt.Println(s_4, &s_4[0], cap(s_4), len(s_4))
	// fmt.Println(s_5, &s_5[0], cap(s_5), len(s_5))
	// /*
	// 	[1 2 3] 0xc00000e198 3 3
	// 	[1 2 3 6] 0xc0000103f0 6 4
	// 	[1 2 3 7] 0xc000010420 6 4
	// */

	// 切片删除元素
	s_6 := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(s_6)
	// s_6 = append(s_6[:1], s_6[4:]...)
	// fmt.Println(s_6)
	copy(s_6[1:], s_6[4:])
	fmt.Println(s_6[:3])
}

func test_map() {
	// map 哈希表
	// 初始化
	m := map[string]int{
		"key1": 1,
		"key2": 2,
	}
	fmt.Println(m)
	// map[key1:1 key2:2]

	m_1 := make(map[string]int)
	m_1["key1"] = 1
	fmt.Println(m_1)
	// map[key1:1]

	delete(m_1, "key1")
	m_2 := make(map[string]int)
	var m_3 map[string]int
	fmt.Println(m_1, m_2, m_3, m_2 == nil, m_3 == nil)
	//map[] map[] map[] false true
}

func test_struct() {
	// struct 结构体
	type name struct {
		a, b  int
		point *name
	}

	p1 := name{}
	p := name{a: 1, b: 2, point: &p1}
	fmt.Println(p)
	// {1 2 0xc000092198}

	// 匿名嵌入测试
	type Struct1 struct {
		name string
	}

	type struct2 struct {
		name int
		Struct1
	}

	test := struct2{11, Struct1{name: "ha"}}
	fmt.Println(test, test.name)
}

func test_async(name string) {
	fmt.Println("it's me ", name)
}

func test_chan() {
	// a := make(chan int, 6)   // 缓冲区为6的双向chan
	// b := make(<-chan int)    // 无缓存区的只读chan
	// c := make(chan<- int, 1) // 缓冲区为1的只写chan

	ch := make(chan int, 1)

	go test_async("jace")

	// 消费者
	go func(ch <-chan int) {
		fmt.Println("wait read")
		result := <-ch
		fmt.Println("from read ", result)
	}(ch)

	// 生产者
	go func(ch chan<- int) {
		time.Sleep(1 * time.Second)
		ch <- 11
	}(ch)

	time.Sleep(3 * time.Second)
	close(ch)

	ch2 := make(chan int)
	fmt.Println(nil == ch2)
}

func main() {

	test_data()

	test_array()

	test_slice()

	test_map()

	test_struct()

	test_chan()
}
