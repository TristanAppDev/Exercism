package flatten

func Flatten(nested interface{}) []interface{} {
	flattened := make([]interface{}, 0)

	nestedPart, ok := nested.([]interface{})
	if !ok {
		return flattened
	}

	for _, n := range nestedPart {
		switch n.(type) {
		case nil:
			continue
		case []interface{}:
			flattened = append(flattened, Flatten(n)...)
		default:
			flattened = append(flattened, n)
		}
	}

	return flattened
}
