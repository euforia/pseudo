package pseudo

import (
	"github.com/hashicorp/hcl/hcl/token"
	"github.com/hashicorp/hil/ast"
)

// VarsMap is a map of variables that may be scoped to a VM
type VarsMap map[string]ast.Variable

// Names returns a list of variable names
func (vars VarsMap) Names() []string {
	out := make([]string, 0, len(vars))
	for k := range vars {
		out = append(out, k)
	}
	return out
}

func variableFromToken(tkn token.Token) ast.Variable {
	v := ast.Variable{Value: tkn.Value()}

	switch tkn.Type {
	case token.NUMBER:
		v.Type = ast.TypeInt

	case token.FLOAT:
		v.Type = ast.TypeFloat

	case token.BOOL:
		v.Type = ast.TypeBool

	default:
		v.Type = ast.TypeString

	}

	return v
}
