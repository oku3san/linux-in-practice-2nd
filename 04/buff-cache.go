package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"time"
)

func main() {

	fmt.Println("最初の free")
	fmt.Println(free())

	var err error
	err = os.Remove("1GB_file.txt")
	if err != nil {
		fmt.Println(err)
	}

	file, _ := os.Create("1GB_file.txt")
	defer file.Close()
	writer := bufio.NewWriter(file)
	defer writer.Flush()
	for i := 0; i < 1024*1024*1024; i++ {
		writer.WriteByte(0)
	}
	fmt.Println("ファイル作成後の free")
	fmt.Println(free())

	err = os.Remove("1GB_file.txt")
	if err != nil {
		fmt.Println(err)
	}

	time.Sleep(time.Second * 5)

	fmt.Println("ファイル削除後の free")
	fmt.Println(free())
}

func free() string {
	result, err := exec.Command("free", "-m").Output()
	if err != nil {
		fmt.Println(err)
	}
	return string(result)
}
