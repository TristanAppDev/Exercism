package robotname

import (
	"math/rand"
)

// Define the Robot type here.
type Robot struct {
	name string
}

var alphabet = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
var robots = make([]Robot, 0)

func (r *Robot) Name() (string, error) {
	if r.name != "" {
		return r.name, nil
	}

	runes := make([]rune, 5)
	for i := 0; i < 5; i++ {
		if i < 2 {
			runes[i] = alphabet[rand.Intn(len(alphabet))]
		} else {
			runes[i] = rune(rand.Intn(10) + '0')
		}
	}
	r.name = string(runes)

	for _, v := range robots {
		if v.name == r.name {
			r.name = ""
			r.Name()
		}
	}

	robots = append(robots, *r)
	return r.name, nil
}

func (r *Robot) Reset() {
	RemoveRobotFromSlice(r)
	r.name = ""
	r.Name()
}

func RemoveRobotFromSlice(r *Robot) {
	for i, v := range robots {
		if v.name == r.name {
			robots = append(robots[:i], robots[i+1:]...)
		}
	}
}
