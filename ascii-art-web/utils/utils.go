package asciiArt

import "bufio"

func InsertValue(scanner *bufio.Scanner) [8]string {
	ArtValue := [8]string{}

	for cp := 0; cp < 8 && scanner.Scan(); cp++ {
		ArtValue[cp] = scanner.Text()
	}
	scanner.Scan()
	return ArtValue
}

func IsValidArg(arg string) bool {
	if arg == "" {
		return false
	}
	for _, val := range arg {
		if val <= 31 || val >= 127 {
			return false
		}
	}
	return true
}

func IsOnlyNewLine(str []string) bool {
	for _, v := range str {
		if v != "" {
			return false
		}
	}
	return true
}
