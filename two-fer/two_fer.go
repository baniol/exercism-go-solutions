package twofer

import "fmt"

// ShareWith returns two-fer string with `you` or a name
// depending on the input
func ShareWith(name string) string {

	x := "you"
	if name != "" {
		x = name
	}

	return fmt.Sprintf("One for %s, one for me.", x)
}
