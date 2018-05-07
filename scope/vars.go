package scope

import (
	"sort"

	"github.com/hashicorp/hil/ast"
)

// Variables holds scoped variables
type Variables map[string]ast.Variable

// Names returns a slice of variable names
func (vars Variables) Names() []string {
	out := make([]string, 0, len(vars))
	for k := range vars {
		out = append(out, k)
	}
	sort.Strings(out)
	return out
}
