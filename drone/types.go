package drone

import (
	"strings"
)

// StringSliceFlag is a flag type which support comma separated values and escaping to not split at unwanted lines
type StringSliceFlag struct {
	slice []string
}

func (s *StringSliceFlag) String() string {
	return strings.Join(s.slice, " ")
}

func (s *StringSliceFlag) Set(value string) error {
	s.slice = splitWithEscaping(value, ",", "\\")
	return nil
}

func (s *StringSliceFlag) Get() []string {
	return s.slice
}

func splitWithEscaping(s, separator, escapeString string) []string {
	if len(s) == 0 {
		return []string{}
	}

	a := strings.Split(s, separator)

	for i := len(a) - 2; i >= 0; i-- {
		if strings.HasSuffix(a[i], escapeString) {
			a[i] = a[i][:len(a[i])-len(escapeString)] + separator + a[i+1]
			a = append(a[:i+1], a[i+2:]...)
		}
	}

	return a
}
