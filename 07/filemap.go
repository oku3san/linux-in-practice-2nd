package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"syscall"
)

func main() {
	pid := os.Getpid()
	fmt.Println(pid)
	fmt.Println("testfileのメモリマップ前のプロセスの仮想アドレス空間")
	command := exec.Command("cat", "/proc/"+strconv.Itoa(pid)+"/maps")
	command.Stdout = os.Stdout
	err := command.Run()
	if err != nil {
		log.Fatal("catの実行に失敗しました")
	}

	file, err := os.OpenFile("testfile", os.O_RDWR, 0)
	if err != nil {
		log.Fatal("testfileを開けませんでした")
	}
	defer file.Close()

	data, err := syscall.Mmap(int(file.Fd()), 0, 5, syscall.PROT_READ|syscall.PROT_WRITE, syscall.MAP_SHARED)
	if err != nil {
		log.Fatal("mmap()に失敗しました")
	}

	fmt.Println("")
	fmt.Printf("testfileをマップしたアドレス: %p\n", &data[0])
	fmt.Println("")

	fmt.Println("testfileのメモリマップ後のプロセスの仮想アドレス空間")
	command = exec.Command("cat", "/proc/"+strconv.Itoa(pid)+"/maps")
	command.Stdout = os.Stdout
	err = command.Run()
	if err != nil {
		log.Fatal("catの実行に失敗しました")
	}

	replaceBytes := []byte("HELLO")
	for i := range data {
		data[i] = replaceBytes[i]
	}
}
