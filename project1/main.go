package main
import (
  "fmt"
  "time"
)
func main() {
    for {
      time.Sleep(time.Duration(2)*time.Second)
      fmt.Println("hello world")
    }
}
