package jsonfieldrename

import (
	"encoding/json"
	"regexp"
)

// TODO: this is super inefficient at the moment
// - switching from bytes to string for the regex then back for the result
// - the regex match includes the separating characters
// - which then have to be stripped out
// - and added back in

// Marshal performs exactly the same task as the std lib marshaller
//, but also renames field keys according to a renaming function that
// you pass in.
func Marshal(entity interface{}, renamer func(string) string) ([]byte, error) {
	b, err := json.Marshal(entity)
	if err != nil {
		return b, err
	}

	s := string(b)
	re := regexp.MustCompile(`"([^"]+)":`)
	new := re.ReplaceAllStringFunc(s, reQuote(renamer))

	return []byte(new), err
}

// reQuote compensates for the fact that I can't work out how to exclude
// the leading " and trailing ": from the regex max
// Firstly, it removes those characters before passing through
// the caller-supplied naming function, then it adds them back in
func reQuote(renamer func(input string) string) func(string) string {
	return func(input string) string {
		return `"` + renamer(input[1:len(input)-2]) + `":`
	}
}
