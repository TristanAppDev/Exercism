package sublist

type Relation string

func Sublist(l1, l2 []int) Relation {
	lenL1 := len(l1)
	lenL2 := len(l2)

	if lenL1 == 0 && lenL2 == 0 {
		return Relation("equal")
	}

	if lenL1 == 0 {
		return Relation("sublist")
	}

	if lenL2 == 0 {
		return Relation("superlist")
	}

	if lenL1 == lenL2 && isEqual(l1, l2, lenL1) {
		return Relation("equal")
	}

	if lenL1 < lenL2 && isSublist(l1, l2, lenL1, lenL2) {
		return Relation("sublist")
	}

	if lenL2 < lenL1 && isSublist(l2, l1, lenL2, lenL1) {
		return Relation("superlist")
	}

	return Relation("unequal")
}

func isSublist(l1, l2 []int, lenL1, lenL2 int) bool {
	for i := 0; i <= lenL2-lenL1; i++ {
		if ok := isEqual(l1, l2[i:i+lenL1], lenL1); ok {
			return true
		}
	}
	return false
}

func isEqual(l1, l2 []int, lenL1 int) bool {
	for j := 0; j < lenL1; j++ {
		if l1[j] != l2[j] {
			return false
		}
	}
	return true
}
