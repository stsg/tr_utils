func translit(r io.Reader, w io.Writer, trans map[rune]rune) error {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := transStr(scanner.Text(), trans)
		if _, err := fmt.Fprintln(w, line); err != nil {
			return err
		}
	}
	return scanner.Err()
}