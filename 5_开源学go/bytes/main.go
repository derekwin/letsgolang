package main

import (
	"bytes"
	"fmt"
)

func split_after() {
	a := []byte("testtesttesttttes")
	result := bytes.SplitAfterN(a, []byte("t"), 2)
	resultS := make([]string, len(result))
	for i := 0; i < len(result); i++ {
		fmt.Println(result[i])
		resultS[i] = string(result[i])
	}
	fmt.Println(resultS)
}

func reader() {
	a := bytes.NewReader([]byte("rsadasdafaadas"))
	fmt.Println(a, a.Len())
	// b := make([]byte, 3)
	// test, _ := a.Read(b)
	// err := a.UnreadByte()
	// fmt.Printf("len:%d, read:%v, b:%v, err:%v", a.Len(), test, string(b), err)
	test, err := a.ReadByte()
	fmt.Println(string(test), err, a, a.Len())
	test, err = a.ReadByte()
	fmt.Println(string(test), err, a, a.Len())
}

func buffer() {
	// a := &bytes.Buffer{}
	// fmt.Println(a, a.Len(), a.Cap())
	// c := append(a.Bytes(), make([]byte, 16)...)
	// fmt.Println(c, len(c), cap(c))

	test := bytes.NewBuffer([]byte("abcdefg"))
	fmt.Println(test, test.Len(), test.Cap(), test.Bytes())
	test.WriteByte(0xc0)
	test.WriteByte(0xc0)
	fmt.Println(test, test.Len(), test.Cap(), test.Bytes())
}

func main() {
	// reader()
	buffer()
}
