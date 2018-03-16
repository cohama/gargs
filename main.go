package main

import (
  "fmt"
  "bufio"
  "os"
  "os/exec"
  "bytes"
)

func main() {
  args := os.Args[1:]
  reader := bufio.NewScanner(os.Stdin)
  for reader.Scan() {
    text := reader.Text()
    cmd := exec.Command(args[0], append(args[1:], text)...)
    result, _ := cmd.Output()
    fmt.Println(text, string(bytes.TrimSpace(result)))
  }
}
