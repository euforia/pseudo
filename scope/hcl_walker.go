package scope

import (
	"fmt"

	cast "github.com/hashicorp/hcl/hcl/ast"
	"github.com/hashicorp/hcl/hcl/token"
	iast "github.com/hashicorp/hil/ast"
)

//
// type hclWalker struct {
// 	prefix string
// 	delim  string
//
// 	vars Variables
// }
//
//
//
// func (w *hclWalker) walk(node cast.Node) map[string]iast.Variable {
// 	out := make(map[string]iast.Variable)
//
// 	switch node.(type) {
//
// 	case *cast.ObjectList:
// 		list := node.(*cast.ObjectList)
// 		for i := range list.Items {
// 			//key := w.prefix
// 			for _, k := range list.Items[i].Keys {
// 				w.prefix += w.delim + k.Token.Text
// 			}
//
// 			items := w.walk(list.Items[i])
// 			for k, v := range items {
// 				out[k] = v
// 			}
// 		}
//
// 	case *cast.ObjectItem:
// 		item := node.(*cast.ObjectItem)
// 		items := w.walkObjectItem(item)
// 		for k, v := range items {
// 			out[k] = v
// 		}
//
// 	case *cast.LiteralType:
// 		tkn := node.(*cast.LiteralType).Token
// 		out[w.prefix] = hilVarTypeFromHCLToken(tkn)
//
// 	}
//
// 	return out
// }
//
// func (w *hclWalker) walkObjectItem(item *cast.ObjectItem) map[string]iast.Variable {
// 	out := make(map[string]iast.Variable)
//
// 	switch item.Val.(type) {
//
// 	case *cast.ObjectType:
// 		ot := item.Val.(*cast.ObjectType)
// 		out = w.walk(ot.List)
//
// 	case *cast.ListType:
// 		lt := item.Val.(*cast.ListType)
// 		l := make([]iast.Variable, len(lt.List))
// 		for i := range lt.List {
// 			//t.prefix = t.prefix + fmt.Sprintf("[%d]", i)
// 			outs := w.walk(lt.List[i])
// 			for k := range outs {
// 				l[i] = outs[k]
// 				break
// 			}
// 		}
// 		out[w.prefix] = iast.Variable{Type: iast.TypeList, Value: l}
//
// 	case *cast.LiteralType:
// 		tkn := item.Val.(*cast.LiteralType).Token
// 		out[w.prefix] = hilVarTypeFromHCLToken(tkn)
//
// 	}
//
// 	return out
// }

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
		out[parent] = hilVarTypeFromHCLToken(tkn)

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
		out[parent] = hilVarTypeFromHCLToken(tkn)

	}

	return out
}

func hilVarTypeFromHCLToken(tkn token.Token) iast.Variable {
	v := iast.Variable{Value: tkn.Value()}

	switch tkn.Type {
	case token.NUMBER:
		v.Type = iast.TypeInt

	case token.FLOAT:
		v.Type = iast.TypeFloat

	case token.BOOL:
		v.Type = iast.TypeBool

	default:
		v.Type = iast.TypeString

	}

	return v
}
