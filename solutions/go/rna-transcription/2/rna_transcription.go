package strand

import "strings"

var transcript = map[rune]rune{
	'G': 'C',
	'C': 'G',
	'T': 'A',
	'A': 'U',
}

func ToRNA(dna string) string {
	var rna strings.Builder
	for _, r := range dna {
		rna.WriteRune(transcript[r])
	}
	return rna.String()
}
