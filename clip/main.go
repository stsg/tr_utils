package main

import (
	"bufio"
	"fmt"
	"io"

	"golang.design/x/clipboard"
)

var trans = map[rune]rune{
	//cyrillic
	'`': 'щ', '-': 'ь', '=': 'ъ',
	'q': 'я', 'w': 'ш', 'e': 'е', 'r': 'р', 't': 'т', 'y': 'ы', 'u': 'у', 'i': 'и', 'o': 'о', 'p': 'п', '[': 'ю', ']': 'ж', '\\': 'э',
	'a': 'а', 's': 'с', 'd': 'д', 'f': 'ф', 'g': 'г', 'h': 'ч', 'j': 'й', 'k': 'к', 'l': 'л',
	'z': 'з', 'x': 'х', 'c': 'ц', 'v': 'в', 'b': 'б', 'n': 'н', 'm': 'м',
	'~': 'Щ', '_': 'Ь', '+': 'Ъ',
	'Q': 'Я', 'W': 'Ш', 'E': 'Е', 'R': 'Р', 'T': 'Т', 'Y': 'Ы', 'U': 'У', 'I': 'И', 'O': 'О', 'P': 'П', '{': 'Ю', '}': 'Ж', '|': 'Э',
	'A': 'А', 'S': 'С', 'D': 'Д', 'F': 'Ф', 'G': 'Г', 'H': 'Ч', 'J': 'Й', 'K': 'К', 'L': 'Л',
	'Z': 'З', 'X': 'Х', 'C': 'Ц', 'V': 'В', 'B': 'Б', 'N': 'Н', 'M': 'М',
	// latin
	'щ': '`', 'ь': '-', 'ъ': '=',
	'я': 'q', 'ш': 'w', 'е': 'e', 'р': 'r', 'т': 't', 'ы': 'y', 'у': 'u', 'и': 'i', 'о': 'o', 'п': 'p', 'ю': '[', 'ж': ']', 'э': '\\',
	'а': 'a', 'с': 's', 'д': 'd', 'ф': 'f', 'г': 'g', 'ч': 'h', 'й': 'j', 'к': 'k', 'л': 'l',
	'з': 'z', 'х': 'x', 'ц': 'c', 'в': 'v', 'б': 'b', 'н': 'n', 'м': 'm',
	'Щ': '~', 'Ь': '_', 'Ъ': '+',
	'Я': 'Q', 'Ш': 'W', 'Е': 'E', 'Р': 'R', 'Т': 'T', 'Ы': 'Y', 'У': 'U', 'И': 'I', 'О': 'O', 'П': 'P', 'Ю': '{', 'Ж': '}', 'Э': '"',
	'А': 'A', 'С': 'S', 'Д': 'D', 'Ф': 'F', 'Г': 'G', 'Ч': 'H', 'Й': 'J', 'К': 'K', 'Л': 'L',
	'З': 'Z', 'Х': 'X', 'Ц': 'C', 'В': 'V', 'Б': 'B', 'Н': 'N', 'М': 'M',
}

func main() {
	err := clipboard.Init()
	if err != nil {
		panic(err)
	}

	newClipboard := []rune(string(clipboard.Read(clipboard.FmtText)))
	newClipboard, err = phonetic_translit(newClipboard)
	if err != nil {
		panic(err)
	}
	clipboard.Write(clipboard.FmtText, []byte(string(newClipboard)))

	// fmt.Println(string(newClipboard))
}

func phonetic_translit(input []rune) ([]rune, error) {
	var outputRunes []rune
	for _, inputChar := range input {
		if outputChar, ok := trans[inputChar]; ok {
			outputRunes = append(outputRunes, outputChar)
		} else {
			outputRunes = append(outputRunes, inputChar)
		}
	}
	return outputRunes, nil
}

func transStr(input string, tr map[rune]rune) string {
	var outputRunes []rune
	for _, inputChar := range input {
		if outputChar, ok := tr[inputChar]; ok {
			outputRunes = append(outputRunes, outputChar)
		} else {
			outputRunes = append(outputRunes, inputChar)
		}
	}
	return string(outputRunes)
}

func translit(r io.Reader, w io.Writer, trans map[rune]rune) error {
	scanner := bufio.NewScanner(bufio.NewReader(r))
	for scanner.Scan() {
		_, e := fmt.Fprintln(
			w, transStr(scanner.Text(), trans))
		if e != nil {
			return e
		}
	}
	return nil
}
