package pseudo

// This file contains scope variable loaders for HCL and HIL format

import (
	"fmt"
	"io/ioutil"

	"github.com/hashicorp/hcl"
	cast "github.com/hashicorp/hcl/hcl/ast"
	iast "github.com/hashicorp/hil/ast"
)

// LoadHCLScopeVarsFromFile loads variables from the given file
func LoadHCLScopeVarsFromFile(filename string) (VarsMap, error) {
	var vars VarsMap

	in, err := ioutil.ReadFile(filename)
	if err == nil {
		vars, err = LoadHCLScopeVars(in)
	}

	return vars, err
}

// LoadHCLScopeVars parses the input to build all the variables that will be in
// this scope
func LoadHCLScopeVars(in []byte) (VarsMap, error) {
	var out VarsMap

	tree, err := hcl.ParseBytes(in)
	if err == nil {
		vars := walk("", tree.Node)
		// Remove the suffix '.'
		out = make(VarsMap, len(vars))
		for k, v := range vars {
			out[k[1:]] = v
		}
	}

	return out, err
}

func walkObjectItem(parent string, item *cast.ObjectItem) map[string]iast.Variable {
	out := make(map[string]iast.Variable)

	switch item.Val.(type) {

	case *cast.ObjectType:
		ot := item.Val.(*cast.ObjectType)
		out = walk(parent, ot.List)

	case *cast.ListType:
		lt := item.Val.(*cast.ListType)
		l := make([]iast.Variable, len(lt.List))
		for i := range lt.List {
			outs := walk(parent+fmt.Sprintf("[%d]", i), lt.List[i])
			for k := range outs {
				l[i] = outs[k]
				break
			}
		}
		out[parent] = iast.Variable{Type: iast.TypeList, Value: l}

	case *cast.LiteralType:
		tkn := item.Val.(*cast.LiteralType).Token
		out[parent] = variableFromToken(tkn)

	}

	return out
}

func walk(parent string, node cast.Node) map[string]iast.Variable {
	out := make(map[string]iast.Variable)

	switch node.(type) {

	case *cast.ObjectList:
		list := node.(*cast.ObjectList)
		for i := range list.Items {
			key := parent
			for _, k := range list.Items[i].Keys {
				key += "." + k.Token.Text
			}

			items := walk(key, list.Items[i])
			for k, v := range items {
				out[k] = v
			}
		}

	case *cast.ObjectItem:
		item := node.(*cast.ObjectItem)
		items := walkObjectItem(parent, item)
		for k, v := range items {
			out[k] = v
		}

	case *cast.LiteralType:
		tkn := node.(*cast.LiteralType).Token
		out[parent] = variableFromToken(tkn)

	}

	return out
}
