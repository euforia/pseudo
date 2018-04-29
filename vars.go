package pseudo

import (
	"io/ioutil"

	"github.com/hashicorp/hcl"
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

// LoadScopeVarsFromFile loads variables from the given file
func LoadScopeVarsFromFile(filename string) (VarsMap, error) {
	var vars VarsMap

	in, err := ioutil.ReadFile(filename)
	if err == nil {
		vars, err = LoadScopeVars(in)
	}

	return vars, err
}

// LoadScopeVars parses the input to build all the variables that will be in
// this scope
func LoadScopeVars(in []byte) (VarsMap, error) {
	var vars VarsMap

	tree, err := hcl.ParseBytes(in)
	if err == nil {
		vars = walk("", tree.Node)
	}

	return vars, err
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
