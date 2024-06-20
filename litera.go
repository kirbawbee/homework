package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Print("Введите текст: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	q := (strings.Count(input, "a"))
	w := (strings.Count(input, "b"))
	e := (strings.Count(input, "c"))
	fmt.Println("a", q, "b", w, "c", e)
}
