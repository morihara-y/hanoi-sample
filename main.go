package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("start hanoi")
	args := os.Args
	n, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Println("err: ", err)
		os.Exit(1)
	}
	cnt := 0
	hanoi(n, args[2], args[3], args[4], &cnt)
	fmt.Println("end hanoi")
}

func hanoi(n int, from, work, dest string, cnt *int) {
	fmt.Printf("call hanoi n: %v, from:%s, work: %s, dest: %s\n", n, from, work, dest)
	if n > 0 {
		hanoi(n-1, from, dest, work, cnt)
		*cnt++
		fmt.Printf("%v, %v番目を、%sから%sへ移動させる。\n", *cnt, n, from, dest)
		hanoi(n-1, work, from, dest, cnt)
	}
}
