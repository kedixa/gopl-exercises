package measure

import (
	"strings"
)

// MethodInefficient Inefficient version
func MethodInefficient(args []string) {
	s, sep := "", ""
	for _, arg := range args {
		s += sep + arg
		sep = " "
	}
}

// MethodJoin shows the strings.Join version
func MethodJoin(args []string) {
	strings.Join(args, " ")
}
