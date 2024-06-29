package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

var trans map[rune]rune = map[rune]rune{
	'`': 'щ', '-': 'ь', '=': 'ъ',
	'q': 'я', 'w': 'ш', 'e': 'е', 'r': 'р', 't': 'т', 'y': 'ы', 'u': 'у', 'i': 'и', 'o': 'о', 'p': 'п', '[': 'ю', ']': 'ж', '\\': 'э',
	'a': 'а', 's': 'с', 'd': 'д', 'f': 'ф', 'g': 'г', 'h': 'ч', 'j': 'й', 'k': 'к', 'l': 'л',
	'z': 'з', 'x': 'х', 'c': 'ц', 'v': 'в', 'b': 'б', 'n': 'н', 'm': 'м',
	'~': 'Щ', '_': 'Ь', '+': 'Ъ',
	'Q': 'Я', 'W': 'Ш', 'E': 'Е', 'R': 'Р', 'T': 'Т', 'Y': 'Ы', 'U': 'У', 'I': 'И', 'O': 'О', 'P': 'П', '{': 'Ю', '}': 'Ж', '|': 'Э',
	'A': 'А', 'S': 'С', 'D': 'Д', 'F': 'Ф', 'G': 'Г', 'H': 'Ч', 'J': 'Й', 'K': 'К', 'L': 'Л',
	'Z': 'З', 'X': 'Х', 'C': 'Ц', 'V': 'В', 'B': 'Б', 'N': 'Н', 'M': 'М',
}

func main() {
	var output []rune
	r := bufio.NewReader(os.Stdin)
	for {
		if c, _, err := r.ReadRune(); err != nil {
			if err == io.EOF {
				break
			} else {
				fmt.Println(err)
				os.Exit(1)
			}
		} else {
			// fmt.Printf("%q [%d]\n", string(c), sz)
			if val, ok := trans[c]; ok {
				output = append(output, val)
				continue
			} else {
				output = append(output, c) // TODO: handle unknown chars
			}
		}
	}

	fmt.Printf("%s\n", string(output))
	return
}
