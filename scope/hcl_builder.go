package scope

import (
	"io/ioutil"

	"github.com/hashicorp/hcl"
)

// BuildHCLScopeVarsFromFile loads variables from the given file
func BuildHCLScopeVarsFromFile(filename string) (Variables, error) {
	var vars Variables

	in, err := ioutil.ReadFile(filename)
	if err == nil {
		vars, err = BuildHCLScopeVars(in)
	}

	return vars, err
}

// BuildHCLScopeVars parses the input to build all the variables that will be in
// this scope
func BuildHCLScopeVars(in []byte) (Variables, error) {
	var out Variables

	tree, err := hcl.ParseBytes(in)
	if err == nil {
		vars := walk("", tree.Node)
		// Remove the suffix '.'
		out = make(Variables, len(vars))
		for k, v := range vars {
			out[k[1:]] = v
		}
	}

	return out, err
}
