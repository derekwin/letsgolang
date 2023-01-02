package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	age := flag.Int("age", 0, "Input age [int]")
	ifAdult := flag.Bool("isAdult", false, "default is false, set true ifAdult")
	var ip string
	flag.StringVar(&ip, "ip", "", "input ip : 192.168.1.1:1080")
	flag.Parse()
	rests := flag.Args()

	test := os.Args
	fmt.Println(*age, ifAdult, ip)
	fmt.Println(test)
	fmt.Println(rests)
}
