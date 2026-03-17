package grains

type InvalidValueError struct{}

func (e InvalidValueError) Error() string {
	return "invalid value"
}

func Square(number int) (uint64, error) {

	if number < 1 || number > 64 {
		return 0, &InvalidValueError{}
	}

	return 1 << (number - 1), nil
}

func Total() uint64 {
	return 1<<64 - 1
}
