package protein

import "errors"

var ErrStop = errors.New("stop sequence found")

var ErrInvalidBase = errors.New("invalid sequence")

func IsStopCodon(c string) bool {
	return c == "UAA" || c == "UAG" || c == "UGA"

}

func FromRNA(rna string) ([]string, error) {
	var proteins []string
	bases := []rune(rna)
	for i := 0; i < len(bases); i += 3 {
		p, err := FromCodon(string(bases[i : i+3]))
		if err == ErrInvalidBase {
			return nil, ErrInvalidBase
		} else if err == ErrStop {
			return proteins, nil
		} else {
			proteins = append(proteins, p)
		}
	}
	return proteins, nil
}

func FromCodon(c string) (string, error) {
	if IsStopCodon(c) {
		return "", ErrStop
	}

	switch c {
	case
		"AUG":
		return "Methionine", nil
	case
		"UUU", "UUC":
		return "Phenylalanine", nil
	case
		"UUA", "UUG":
		return "Leucine", nil
	case
		"UCU", "UCC", "UCA", "UCG":
		return "Serine", nil
	case
		"UAU", "UAC":
		return "Tyrosine", nil
	case
		"UGU", "UGC":
		return "Cysteine", nil
	case
		"UGG":
		return "Tryptophan", nil
	default:
		return "", ErrInvalidBase
	}
}
