package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

var trans_phonetic_mac = map[rune]rune{
	// cyrillic
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

var trans_phonetic_winkeys = map[rune]rune{
	// cyrillic
	'`': 'ю', '#': 'ё', '$': 'Ё', '%': 'ъ', '^': 'Ъ', '=': 'ь',
	'q': 'я', 'w': 'в', 'e': 'е', 'r': 'р', 't': 'т', 'y': 'ы', 'u': 'у', 'i': 'и', 'o': 'о', 'p': 'п', '[': 'ш', ']': 'щ', '\\': 'э',
	'a': 'а', 's': 'с', 'd': 'д', 'f': 'ф', 'g': 'г', 'h': 'ч', 'j': 'й', 'k': 'к', 'l': 'л',
	'z': 'з', 'x': 'х', 'c': 'ц', 'v': 'ж', 'b': 'б', 'n': 'н', 'm': 'м',
	'~': 'Ю', '+': 'Ь',
	'Q': 'Я', 'W': 'В', 'E': 'Е', 'R': 'Р', 'T': 'Т', 'Y': 'Ы', 'U': 'У', 'I': 'И', 'O': 'О', 'P': 'П', '{': 'Ш', '}': 'Щ', '|': 'Э',
	'A': 'А', 'S': 'С', 'D': 'Д', 'F': 'Ф', 'G': 'Г', 'H': 'Ч', 'J': 'Й', 'K': 'К', 'L': 'Л',
	'Z': 'З', 'X': 'Х', 'C': 'Ц', 'V': 'Ж', 'B': 'Б', 'N': 'Н', 'M': 'М',
	// latin
	'ю': '`', 'ё': '#', 'Ё': '$', 'ъ': '%', 'Ъ': '^', 'ь': '=',
	'я': 'q', 'в': 'w', 'е': 'e', 'р': 'r', 'т': 't', 'ы': 'y', 'у': 'u', 'и': 'i', 'о': 'o', 'п': 'p', 'ш': '[', 'щ': ']', 'э': '\\',
	'а': 'a', 'с': 's', 'д': 'd', 'ф': 'f', 'г': 'g', 'ч': 'h', 'й': 'j', 'к': 'k', 'л': 'l',
	'з': 'z', 'х': 'x', 'ц': 'c', 'ж': 'v', 'б': 'b', 'н': 'n', 'м': 'm',
	'Ю': '~', 'Ь': '+',
	'Я': 'Q', 'В': 'W', 'Е': 'E', 'Р': 'R', 'Т': 'T', 'Ы': 'Y', 'У': 'U', 'И': 'I', 'О': 'O', 'П': 'P', 'Ш': '{', 'Щ': '}', 'Э': '"',
	'А': 'A', 'С': 'S', 'Д': 'D', 'Ф': 'F', 'Г': 'G', 'Ч': 'H', 'Й': 'J', 'К': 'K', 'Л': 'L',
	'З': 'Z', 'Х': 'X', 'Ц': 'C', 'Ж': 'V', 'Б': 'B', 'Н': 'N', 'М': 'M',
}

var revision = ""

// main is the entry point of the program.
//
// It reads input from the standard input, translates it using the `trans` map,
// and writes the translated output to the standard output.
//
// It returns an error if there is any issue during the translation process.
func main() {
	err := translit(os.Stdin, os.Stdout, trans_phonetic_winkeys)
	if err != nil {
		panic(err)
	}
}

// transStr transcribes a string using a given map of rune translations.
//
// The function takes an input string and a map of rune translations as parameters.
// It iterates over each rune in the input string and checks if it exists in the
// translation map. If a translation is found, the corresponding rune is appended
// to the outputRunes slice. If no translation is found, the original rune is
// appended to the outputRunes slice.
//
// The function returns the transcribed string as a string.
func transStr(input string, tr map[rune]rune) string {
	outputRunes := make([]rune, 0, len(input))
	for _, inputChar := range input {
		if outputChar, ok := tr[inputChar]; ok {
			outputRunes = append(outputRunes, outputChar)
		} else {
			outputRunes = append(outputRunes, inputChar)
		}
	}
	return string(outputRunes)
}

// translit transcribes text read from r using the translation map trans and writes the transcribed text to w.
//
// Parameters:
//   - r: an io.Reader from which text is read.
//   - w: an io.Writer to which transcribed text is written.
//   - trans: a map of rune translations used for transliteration.
//
// Returns an error if any.
func translit(r io.Reader, w io.Writer, trans map[rune]rune) error {
	scanner := bufio.NewScanner(r)
	var output []string
	for scanner.Scan() {
		output = append(output, transStr(scanner.Text(), trans))
	}
	last := len(output) - 1
	if last < 0 {
		return nil
	}
	for i := 0; i < last; i++ {
		_, err := fmt.Fprintln(w, output[i])
		if err != nil {
			return err
		}
	}
	_, err := fmt.Fprint(w, output[last])
	if err != nil {
		return err
	}

	return nil
}
