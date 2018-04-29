package pseudo

import (
	"fmt"

	cast "github.com/hashicorp/hcl/hcl/ast"
	iast "github.com/hashicorp/hil/ast"
)

func walkItem(parent string, item *cast.ObjectItem) map[string]iast.Variable {
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
		items := walkItem(parent, item)
		for k, v := range items {
			out[k] = v
		}

	case *cast.LiteralType:
		tkn := node.(*cast.LiteralType).Token
		out[parent] = variableFromToken(tkn)

	}

	return out
}
