package main

import (
  "fmt"
  "os"
  "os/signal"
  "syscall"
)

func main() {
  sigs := make(chan os.Signal, 1)
  signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
  signal.Ignore(os.Interrupt)

  go func() {
    s := <-sigs
    fmt.Println("Got signal:", s)
  }()
  fmt.Println("loop start")
  for {
  }
}
