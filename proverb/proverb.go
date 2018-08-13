package proverb

import "fmt"

// Proverb returns a list of proverbs, generated from the input words.
// For example, provided the input `[]string{"nail", "shoe", "horse"}`, the output list must contain:
// "For want of a nail the shoe was lost.", "For want of a shoe the horse was lost.", "And all for the want of a nail."
func Proverb(rhyme []string) []string {

	rhymeLenght := len(rhyme)

	if rhymeLenght == 0 {
		return []string{}
	}

	out := make([]string, rhymeLenght)

	for i := 0; i < rhymeLenght-1; i++ {
		out[i] = fmt.Sprintf("For want of a %s the %s was lost.", rhyme[i], rhyme[i+1])
	}

	out[rhymeLenght-1] = fmt.Sprintf("And all for the want of a %s.", rhyme[0])

	return out
}
