// Package twofer is a exercise in Go from exercism.org
package twofer

// ShareWith returns a string with the name of the person passed as parameter
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return "One for " + name + ", one for me."
}
