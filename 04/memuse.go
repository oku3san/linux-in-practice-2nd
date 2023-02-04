package main

import (
	"fmt"
	"os/exec"
)

var size int64 = 10000000

func main() {
	result, err := exec.Command("free", "-m").Output()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Print(string(result))

	slice := make([]uint8, size)
	fmt.Println(slice)

	result2, err2 := exec.Command("free", "-m").Output()
	if err2 != nil {
		fmt.Println(err2)
	}
	fmt.Print(string(result2))

}
