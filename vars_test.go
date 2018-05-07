package pseudo

// func Test_BuildVarScope(t *testing.T) {
// 	m, err := LoadHCLScopeVarsFromFile(testScopeVarsSpec)
// 	assert.Nil(t, err)
// 	t.Log(m.Names())
// }

// func Test_walk(t *testing.T) {
// 	in := []byte(`
//         number = 4
//         float = 1.1
//         bool = true
// `)
// 	tree, _ := hcl.ParseBytes(in)
// 	vars := walk("", tree.Node)
//
// 	v, _ := vars[".number"]
// 	assert.Equal(t, ast.TypeInt, v.Type)
//
// 	v, _ = vars[".float"]
// 	assert.Equal(t, ast.TypeFloat, v.Type)
//
// 	v, _ = vars[".bool"]
// 	assert.True(t, v.Value.(bool))
// }

//
// func Test_scalarVarFromReflectValue(t *testing.T) {
//
// 	opt := IndexOptions{ContentType: "hcl"}
// 	uri, _ := url.Parse(testPlatformSpec)
//
// 	idx, err := LoadIndex(uri, opt)
// 	assert.Nil(t, err)
//
// 	m := make(map[string]ast.Variable)
// 	idx.Iter(func(key string, value reflect.Value) bool {
// 		m[key] = scalarVarFromReflectValue(value)
// 		return true
// 	})
//
// }
