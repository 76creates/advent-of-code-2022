package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, _ := os.Open("../input")
	defer f.Close()

	s := bufio.NewScanner(f)
	s.Scan()
	line := s.Text()

	var buffer []string
	for position, char := range line {
		if len(buffer) < 14 {
			buffer = append(buffer, string(char))
		} else {
			buffer = buffer[1:]
			buffer = append(buffer, string(char))
		}
		if len(buffer) == 14 {
			unique := make(map[string]struct{})
			for i := 0; i < len(buffer); i++ {
				unique[buffer[i]] = struct{}{}
			}
			if len(unique) == 14 {
				fmt.Println(position + 1)
				return
			}
		}
	}
}
