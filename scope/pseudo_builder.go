package scope

import (
	"github.com/hashicorp/hcl"
)

type PseudoBuilder struct {
	prefix string
	vars   Variables
}

func NewPseudoBuilder(prefix string) *PseudoBuilder {
	return &PseudoBuilder{prefix: prefix, vars: make(Variables)}
}

func (builder *PseudoBuilder) Build(in []byte) error {
	tree, err := hcl.ParseBytes(in)
	if err == nil {
		vars := walk(builder.prefix, tree.Node)
		// Remove the suffix '.'
		builder.vars = make(Variables, len(vars))
		for k, v := range vars {
			builder.vars[k[1:]] = v
		}
	}

	return err
}

// Variables returns the all available variables
func (builder *PseudoBuilder) Variables() Variables {
	return builder.vars
}

// BuildHCLScopeVarsFromFile loads variables from the given file
// func BuildHCLScopeVarsFromFile(filename string) (Variables, error) {
// 	var vars Variables
//
// 	in, err := ioutil.ReadFile(filename)
// 	if err == nil {
// 		//vars, err = BuildHCLScopeVars(in)
// 		b := NewPseudoBuilder("")
// 		if err = b.Build(in); err == nil {
// 			vars = b.Variables()
// 		}
// 	}
//
// 	return vars, err
// }

// BuildHCLScopeVars parses the input to build all the variables that will be in
// this scope
// func BuildHCLScopeVars(in []byte) (Variables, error) {
// 	var out Variables
//
// 	tree, err := hcl.ParseBytes(in)
// 	if err == nil {
// 		vars := walk("", tree.Node)
// 		// Remove the suffix '.'
// 		out = make(Variables, len(vars))
// 		for k, v := range vars {
// 			out[k[1:]] = v
// 		}
// 	}
//
// 	return out, err
// }
