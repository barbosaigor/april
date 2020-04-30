package selector

import (
	"fmt"
	"regexp"
)

// Selector represents all possible match functions
var Selector = map[string]int{
	"prefix":  Prefix,
	"infix":   Infix,
	"postfix": Postfix,
	"all":     All,
}

// Possible states for matching operations
const (
	Prefix = iota
	Infix
	Postfix
	All
)

func generateExp(pattern string, notation int) string {
	var exp string
	switch notation {
	case All:
		exp = `^%s$`
	case Prefix:
		exp = `^%s`
	case Infix:
		exp = `%s`
	case Postfix:
		exp = `%s$`
	}
	return fmt.Sprintf(exp, regexp.QuoteMeta(pattern))
}

func createRegex(pattern string, notation int) *regexp.Regexp {
	return regexp.MustCompile(generateExp(pattern, notation))
}

// Match finds a pattern using a notation (Prefix, Infix, Postfix or All pattern match)
// if the pattern was found it'll return true, elsewere return false.
// If it's invalid notation then panic.
func Match(txt, pattern string, notation int) bool {
	if notation != Prefix && notation != Infix && notation != Postfix && notation != All {
		panic("Invalid notation")
	}
	regexpr := createRegex(pattern, notation)
	return regexpr.MatchString(txt)
}
