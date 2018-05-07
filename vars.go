package pseudo

//
// // VarsMap is a map of variables that may be scoped to a VM
// type VarsMap map[string]ast.Variable
//
// // Names returns a list of variable names
// func (vars VarsMap) Names() []string {
// 	out := make([]string, 0, len(vars))
// 	for k := range vars {
// 		out = append(out, k)
// 	}
// 	return out
// }

// func scalarVarFromHCLToken(tkn token.Token) ast.Variable {
// 	v := ast.Variable{Value: tkn.Value()}
//
// 	switch tkn.Type {
// 	case token.NUMBER:
// 		v.Type = ast.TypeInt
//
// 	case token.FLOAT:
// 		v.Type = ast.TypeFloat
//
// 	case token.BOOL:
// 		v.Type = ast.TypeBool
//
// 	default:
// 		v.Type = ast.TypeString
//
// 	}
//
// 	return v
// }

//
// func scalarVarFromReflectValue(v reflect.Value) ast.Variable {
// 	avar := ast.Variable{Value: v.Interface()}
//
// 	kind := v.Kind()
//
// 	switch {
// 	case kind >= reflect.Int && kind <= reflect.Uintptr:
// 		avar.Type = ast.TypeInt
//
// 	case kind >= reflect.Float32 && kind <= reflect.Complex128:
// 		avar.Type = ast.TypeFloat
//
// 	case kind == reflect.Bool:
// 		avar.Type = ast.TypeBool
//
// 	case kind == reflect.String:
// 		avar.Type = ast.TypeString
//
// 	}
//
// 	return avar
// }
