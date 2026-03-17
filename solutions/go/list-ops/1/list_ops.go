package listops

// IntList is an abstraction of a list of integers which we can define methods on
type IntList []int

func (s IntList) Foldl(fn func(int, int) int, initial int) int {
	if len(s) == 0 {
		return initial
	}
	x, xs := s[0], s[1:]
	return xs.Foldl(fn, fn(initial, x))
}

func (s IntList) Foldr(fn func(int, int) int, initial int) int {
	reversedFunc := func(x, y int) int { return fn(y, x) }
	return s.Reverse().Foldl(reversedFunc, initial)
}

func (s IntList) Filter(fn func(int) bool) IntList {
	var result = []int{}
	for _, v := range s {
		if fn(v) {
			result = append(result, v)
		}
	}
	return result
}

func (s IntList) Length() int {
	return len(s)
}

func (s IntList) Map(fn func(int) int) IntList {
	var result = make([]int, s.Length())
	for i, v := range s {
		result[i] = fn(v)
	}
	return result
}

func (s IntList) Reverse() IntList {
	var result = make([]int, s.Length())
	for i, j := 0, s.Length()-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = s[j], s[i]
	}
	return result
}

func (s IntList) Append(lst IntList) IntList {
	return append(s, lst...)
}

func (s IntList) Concat(lists []IntList) IntList {

	var result = make([]int, 0, s.Length())

	result = append(result, s...)

	for _, l := range lists {
		if l.Length() > 0 {
			result = append(result, l...)
		}
	}
	return result
}
