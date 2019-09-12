package main

import (
    "fmt"
    "testing"
    "time"
)

func main() {
    result := testing.Benchmark(func(b *testing.B) { run() })
    fmt.Println(result.T)
}

func run() {
    fmt.Println("Start!")
    go process("A")
    go process("B")
    go process("C")
    time.Sleep(2 * time.Second)
    fmt.Println("Finish!")
}

func process(name string) {
    time.Sleep(2 * time.Second)
    fmt.Println(name)
}
