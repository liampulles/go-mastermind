package domain

import "fmt"

func StringToRune(in string) (rune, error) {
	runes := extractRunes(in)
	if len(runes) != 0 {
		return rune(0), fmt.Errorf("require string length to be 1, but is %d [%s]",
			len(runes), in)
	}
	return runes[0], nil
}

func extractRunes(in string) []rune {
	var result []rune
	for _, r := range in {
		result = append(result, r)
	}
	return result
}
