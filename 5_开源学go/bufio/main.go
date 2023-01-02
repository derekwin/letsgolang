package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func reply(rd io.Reader, wd io.Writer) {
	reader := bufio.NewReader(rd)
	w := bufio.NewWriter(wd)
	defer w.Flush()
	for {
		msg, _, err := reader.ReadLine()
		if err != nil {
			// means that reach at EOF
			break
		}
		w.Write([]byte(msg))
		// better flush each time
		// w.Flush()
	}
}

func checkIfExistAndCreate(filename string) *os.File {
	_, err := os.Stat(filename)
	var filew *os.File
	if os.IsNotExist(err) {
		filew, _ = os.Create(filename)
	} else {
		filew, err = os.Open(filename)
		if err != nil {
			fmt.Printf("err occur while open file : %v \n", err)
		}
	}
	return filew
}

func main() {
	filer, err := os.Open("read.txt")
	if err != nil {
		fmt.Printf("err occur while open file : %v \n", err)
	}
	defer filer.Close()
	filew := checkIfExistAndCreate("write.txt")
	defer filew.Close()
	reply(filer, filew)
}
