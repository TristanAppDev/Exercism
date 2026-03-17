// twofer is a exercise in Go from exercism.org
package twofer

import "fmt"

// ShareWith returns a string with the name of the person passed as parameter
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}
