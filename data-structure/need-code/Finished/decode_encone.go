package main

import (
	"fmt"
	"strings"
	"strconv"
)



type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	var b strings.Builder
	for _, x := range strs {
		L := len(x)
		stringValue := fmt.Sprintf("%d#%s", L, x)
		b.WriteString(stringValue)
	}
	return b.String()
}


func (s *Solution) Decode(encoded string) []string {
	var decoded []string
	i := 0

	for i < len(encoded) {
		j := i
		for j < len(encoded) && encoded[j] != '#' {
			j++
		}
		lengthText := encoded[i:j]
		i = j + 1
		n, err := strconv.Atoi(lengthText)
		if err != nil {
			fmt.Println("Bad", lengthText, err)
			return decoded
		}
		payload := encoded[i : i+n]
		decoded = append(decoded,payload)
		i += n
	}
	return decoded
}


func main() {
	s := &Solution{}

	// Example placeholders so you can test as you build:
	input := []string{"hello", "world"}
	encoded := s.Encode(input)
	decoded := s.Decode(encoded)

	fmt.Println("Encoded:", encoded)
	fmt.Println("Decoded:", decoded)
}
